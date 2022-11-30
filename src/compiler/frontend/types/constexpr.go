package types

func EvalConstBoolBinOp(op string, t1, t2 Type) *bool {
	cv1, ok1 := t1.Const()
	cv2, ok2 := t2.Const()
	if !ok1 || !ok2 {
		return nil
	}

	var ret bool
	switch cv1 := cv1.(type) {
	case (int):
		cv2 := cv2.(int)
		switch op {
		case "<":
			ret = cv1 < cv2
		case ">":
			ret = cv1 > cv2
		case "<=":
			ret = cv1 <= cv2
		case ">=":
			ret = cv1 >= cv2
		case "==":
			ret = cv1 == cv2
		case "!=":
			ret = cv1 != cv2
		default:
			panic("unexpected bool bin op")
		}
	case (bool):
		cv2 := cv2.(bool)
		switch op {
		case "&&":
			ret = cv1 && cv2
		case "||":
			ret = cv1 || cv2
		case "==":
			ret = cv1 == cv2
		case "!=":
			ret = cv1 != cv2
		default:
			panic("unexpected bool bin op")
		}
	case (string):
		cv2 := cv2.(string)
		switch op {
		case "<":
			ret = cv1 < cv2
		case ">":
			ret = cv1 > cv2
		case "<=":
			ret = cv1 <= cv2
		case ">=":
			ret = cv1 >= cv2
		case "==":
			ret = cv1 == cv2
		case "!=":
			ret = cv1 != cv2
		default:
			panic("unexpected bool bin op")
		}

	default:
		return nil
	}

	return &ret
}

func EvalConstIntBinOp(op string, t1, t2 Type) *int {
	cv1, ok1 := t1.Const()
	cv2, ok2 := t2.Const()

	if ok1 && ok2 {
		var ret int
		cv1 := cv1.(int)
		cv2 := cv2.(int)
		switch op {
		case "-":
			ret = cv1 - cv2
		case "+":
			ret = cv1 + cv2
		case "*":
			ret = cv1 * cv2
		case "/":
			ret = cv1 / cv2
		case "%":
			ret = cv1 % cv2
		default:
			panic("unexpected int bin op")
		}
		return &ret
	}

	return nil
}

func EvalConstStringBinOp(op string, t1, t2 Type) *string {
	cv1, ok1 := t1.Const()
	cv2, ok2 := t2.Const()

	if ok1 && ok2 {
		var ret string
		cv1 := cv1.(string)
		cv2 := cv2.(string)
		switch op {
		case "+":
			ret = cv1 + cv2
		default:
			panic("unexpected string bin op")
		}
		return &ret
	}

	return nil
}

func EvalConstIntNegOp(t TInt) *int {
	if t.Value == nil {
		return nil
	}

	res := -*t.Value
	return &res
}

func EvalConstBoolNotOp(t TBool) *bool {
	if t.Value == nil {
		return nil
	}

	res := !*t.Value
	return &res
}
