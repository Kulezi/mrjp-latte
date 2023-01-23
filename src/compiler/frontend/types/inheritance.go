package types

func (s *Signatures) inheritClass(child string, parent TClassRef, evaluated map[string]struct{}) error {
	// Resolve parent's inheritance if it's a derived class.
	if _, ok := s.Globals[parent.String()]; !ok {
		return UnknownClassError{
			Type: parent,
		}
	}

	if _, ok := evaluated[parent.String()]; !ok {
		if grandparent, ok := s.Parent[parent.String()]; ok {
			err := s.inheritClass(parent.String(), grandparent, evaluated)
			if err != nil {
				return err
			}
		}
	}

	childClass := s.Globals[child].Type.(TClass)
	childFields := childClass.Fields
	parentClass := s.Globals[parent.String()].Type.(TClass)
	parentFields := parentClass.Fields

	// Inherit methods.
	for ident, parentFieldInfo := range parentFields {
		parentFieldType, ok := parentFieldInfo.Type.(TFun)
		if !ok {
			continue
		}

		newFieldInfo := parentFieldInfo
		// Function overriding needs matching signatures.
		if childFieldInfo, ok := childFields[ident]; ok {
			if !SameType(parentFieldType, childFieldInfo.Type) {
				return MethodOverrideError{
					ParentClass:  s.Globals[parent.String()],
					ChildClass:   s.Globals[child],
					ParentMethod: parentFieldType,
					ChildMethod:  childFieldInfo.Type,
					MethodName:   ident,
				}
			}
			newFieldInfo.Origin = childFieldInfo.Origin
			newFieldInfo.Type = childFieldInfo.Type
		}

		// Overriden methods have the same offset in vtable.
		childFields[ident] = newFieldInfo
	}

	// Eval offsets for fields and methods that are not present in the parent.
	childClass.TotalMethods = parentClass.TotalMethods
	childClass.TotalNonMethods = parentClass.TotalNonMethods
	for ident, childFieldInfo := range childFields {
		switch childFieldInfo.Type.(type) {
		case TFun:
			if _, ok := parentFields[ident]; !ok {
				childFieldInfo.Offset = childClass.TotalMethods
				childClass.TotalMethods++
			}
		default:
			childClass.TotalNonMethods++
			childFieldInfo.Offset = childClass.TotalNonMethods
		}

		childFields[ident] = childFieldInfo
	}
	evaluated[child] = struct{}{}

	s.ReplaceGlobal(child, childClass)
	return nil
}

func (s *Signatures) InheritClasses() error {
	evaluated := make(map[string]struct{})
	for class, parent := range s.Parent {
		if _, ok := evaluated[class]; ok {
			continue
		}

		var err error
		if err = s.inheritClass(class, parent, evaluated); err != nil {
			return err
		}
	}

	return nil
}
