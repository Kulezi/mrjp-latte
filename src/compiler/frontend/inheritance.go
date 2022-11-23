package frontend

func (s *Signatures) inheritClass(child string, parent TClassRef, evaluated map[string]struct{}) (TClass, error) {
	// Resolve parent's inheritance if it's a derived class.
	var err error
	if _, ok := evaluated[parent.String()]; !ok {
		if grandparent, ok := s.Parent[parent.String()]; ok {
			s.Globals[parent.String()], err = s.inheritClass(parent.String(), grandparent, evaluated)
			if err != nil {
				return TClass{}, err
			}
		}
	}

	childClass := s.Globals[child].(TClass)
	childFields := childClass.Fields
	parentFields := s.Globals[parent.String()].(TClass).Fields
	for ident, parentFieldType := range parentFields {
		if childFieldType, ok := childFields[ident]; ok {
			parentFieldType, ok := parentFieldType.(TFun)
			if !ok {
				continue
			}

			if !sameType(parentFieldType, childFieldType) {
				return TClass{}, MethodOverrideError{
					ParentClass:  s.Globals[parent.String()],
					ChildClass:   s.Globals[child],
					ParentMethod: parentFieldType,
					ChildMethod:  childFieldType,
					MethodName:   ident,
				}
			}

			childFields[ident] = parentFieldType
		}
	}

	evaluated[child] = struct{}{}
	return childClass, nil
}

func (s *Signatures) inheritClasses() error {
	evaluated := make(map[string]struct{})
	for class, parent := range s.Parent {
		if _, ok := evaluated[class]; ok {
			continue
		}

		var err error
		if s.Globals[class], err = s.inheritClass(class, parent, evaluated); err != nil {
			return err
		}
	}

	return nil
}
