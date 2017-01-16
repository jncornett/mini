package mini

// Op is the index of a builtin operation to invoke on an Object
type Op int

const (
	OpNoop Op = iota
	OpEq
	OpNe
	OpNeg
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpLt
	OpLe
	OpGt
	OpGe
)

func (o Op) String() string {
	switch o {
	case OpNoop:
		return "noop"
	case OpEq:
		return "eq"
	case OpNe:
		return "ne"
	case OpNeg:
		return "neg"
	case OpAdd:
		return "add"
	case OpSub:
		return "sub"
	case OpMul:
		return "mul"
	case OpDiv:
		return "div"
	case OpLt:
		return "lt"
	case OpLe:
		return "le"
	case OpGt:
		return "gt"
	case OpGe:
		return "ge"
	}
	return "?"
}
