package sym

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	lltyp "github.com/llir/llvm/ir/types"
	llval "github.com/llir/llvm/ir/value"
	"log"
	"regexp"
	"strconv"
)

var EExist =fmt.Errorf("symbol already exist")
var ENotSupport =fmt.Errorf("not supported op")
var ESymbolTypeError =fmt.Errorf("symbol type error")
var ETupleTooLong =fmt.Errorf("tuple too long > 10")
var EInternal =fmt.Errorf("internal error")

type Sym interface {
	Type() symType
}

type SymTab interface {
	Resolve(name string) Sym
	Parent() SymTab
	Add(name string, sym Sym) error
	SetParent(s SymTab)
}

type symType int

const (
	SFunction symType = iota
	SValue
	SType
)



type ParentSymTab struct {
	parent SymTab
}

func (p *ParentSymTab) Parent() SymTab {
	return p.parent
}

func (p *ParentSymTab) Add(name string, sym Sym) error {
	return ENotSupport
}

func (p *ParentSymTab) SetParent(t SymTab) {
	p.parent=t
}

type SymbolTable struct {
	ParentSymTab
	dict   map[string]Sym
}

func (s SymbolTable) Resolve(name string) Sym {
	if r,ok:=s.dict[name];ok{
		return r
	}
	if s.parent !=nil{
		return s.parent.Resolve(name)
	}
	return nil
}

func (s *SymbolTable) Add(name string, sym Sym) error{
	if _,ok:=s.dict[name];ok{
		return EExist
	}
	s.dict[name]=sym
	return nil
}

func NewSymTable(parent SymTab) SymTab {
	return &SymbolTable{ParentSymTab: ParentSymTab{parent:parent}, dict: map[string]Sym{}}
}

func NewBuiltinTable() SymTab {
	st:=NewSymTable(nil)
	_ = st.Add("int", IntTyp)
	_ = st.Add("float", FloatTyp)
	_ = st.Add("bool", BoolTyp)
	_ = st.Add("true", TrueVal)
	_ = st.Add("false", FalseVal)
	_ = st.Add("string", StringTyp)
	_ = st.Add(UnitPhonyName, UnitTyp)
	return st
}

type FuncTyp struct {
	Arg TypSym
	Ret TypSym
	LlType *lltyp.FuncType
}

var _ TypSym = &FuncTyp{}

func (f *FuncTyp) LlTyp() lltyp.Type {
	return f.LlType
}

func (f *FuncTyp) NewInstance() ValSym {
	v:=&FuncVal{typ: f, ArgVal: f.Arg.NewInstance().(*TupleVal)}
	return v
}

func (f *FuncTyp) Type() symType {
	return SType
}

func GetFuncTyp(arg TypSym, ret TypSym)*FuncTyp{
	t :=[2]TypSym{arg, ret}
	if _,ok:=funcTypePool[t];!ok{
		tp:=&FuncTyp{Arg: arg, Ret: ret}
		funcTypePool[t] = tp
	}
	return funcTypePool[t]
}

var funcTypePool = map[[2]TypSym]*FuncTyp{}

type FuncVal struct {
	llVal llval.Value
	typ   *FuncTyp
	argNames []string
	ArgVal ValSym
	RetVal ValSym
	ParentSymTab
}

func (f *FuncVal) SetLlVal(value llval.Value) error {
	f.llVal =value
	return nil
}

var _ SymTab = &FuncVal{}

func (f *FuncVal) Resolve(name string) Sym {
	idx:=-1
	for i, n := range f.argNames {
		if n==name{
			idx=i
			break
		}
	}
	if idx==-1{
		return f.parent.Resolve(name)
	}
	if len(f.argNames)==1{
		if idx!=0{
			panic(EInternal)
		}
		return f.ArgVal
	}

	return f.ArgVal.(*TupleVal).Members[idx]
}

func (f *FuncVal) Type() symType {
	return SValue
}

func (f *FuncVal) LlVal() llval.Value {
	return f.llVal
}

func (f *FuncVal) Typ() TypSym {
	return f.typ
}

func (f *FuncVal) AutoFill(){
	f.ArgVal = f.ArgVal.(*TupleVal).Reduce()
	f.typ=GetFuncTyp(f.ArgVal.Typ(), f.RetVal.Typ())
}

func (f *FuncVal) Add(name string, sym Sym) error{
	if f.Resolve(name)!=nil{
		return EExist
	}
	f.argNames = append(f.argNames, name)
	if sym.Type()!=SType{
		return ESymbolTypeError
	}
	nv:=sym.(TypSym).NewInstance()
	f.ArgVal.(*TupleVal).Members = append(f.ArgVal.(*TupleVal).Members, nv)
	return nil
}

type TypSym interface {
	Sym
	LlTyp() lltyp.Type
	NewInstance() ValSym
}

type PrimaryTyp struct {
	LlType lltyp.Type
}

func (t *PrimaryTyp) NewInstance() ValSym {
	switch t {
	case IntTyp:
		fallthrough
	case ByteTyp:
		fallthrough
	case CharTyp:
		fallthrough
	case FloatTyp:
		fallthrough
	case StringTyp:
		fallthrough
	case BoolTyp:
		return &PrimaryVal{llVal: nil, typ: t}
	case UnitTyp:
		return UnitVal
	}
	return nil
}

var _ TypSym = &PrimaryTyp{}

func (t *PrimaryTyp) Type() symType {
	return SType
}

func (t *PrimaryTyp) LlTyp() lltyp.Type {
	return t.LlType
}

var (
	UnitTyp = &PrimaryTyp{LlType: lltyp.Void}
	BoolTyp = &PrimaryTyp{LlType: lltyp.I1}
	IntTyp = &PrimaryTyp{LlType: lltyp.I64}
	ByteTyp = &PrimaryTyp{LlType: lltyp.I8}
	CharTyp = &PrimaryTyp{LlType: lltyp.NewInt(8)}
	StringTyp = &PrimaryTyp{LlType: lltyp.I8Ptr}
	FloatTyp = &PrimaryTyp{LlType: lltyp.Double}
)

type TupleTyp struct {
	llTyp   lltyp.Type
	members []TypSym
}

var _ TypSym=&TupleTyp{}

func (t *TupleTyp) Type() symType {
	return SType
}

func (t *TupleTyp) NewInstance() ValSym {
	members:=[]ValSym{}
	for _, m := range t.members {
		members = append(members, m.NewInstance())
	}
	v:=&TupleVal{typ: t, Members: members}
	return v
}

func (t *TupleTyp) LlTyp() lltyp.Type {
	if t.llTyp==nil{
		var llTypes []lltyp.Type
		for _, m := range t.members {
			tp:=m.LlTyp()
			if tp==nil{
				return nil
			}
			llTypes = append(llTypes, tp)
		}
		t.llTyp = lltyp.NewStruct(llTypes...)
	}
	return t.llTyp
}

type NameTypPair struct {
	Name string
	Typ TypSym
}

type StructTyp struct {
	llTyp   lltyp.Type
	members []NameTypPair
	parent SymTab
}

var _ TypSym = &StructTyp{}

func (t *StructTyp) NewInstance() ValSym {
	members:=[]ValSym{}
	for _, m := range t.members {
		members = append(members, m.Typ.NewInstance())
	}
	v:=&StructVal{typ: t, members: members}
	return v
}

func (t *StructTyp) Resolve(name string) Sym {
	for _, m := range t.members {
		if m.Name == name{
			return m.Typ
		}
	}

	if t.parent!=nil{
		return t.parent.Resolve(name)
	}
	return nil
}

func (t *StructTyp) Parent() SymTab {
	return t.parent
}

func (t *StructTyp) Fill() error {
	fields := make([]lltyp.Type,0, len(t.members))
	for _, member := range t.members {
		fields = append(fields, member.Typ.LlTyp())
	}
	t.llTyp= lltyp.NewStruct(fields...)
	return nil
}

func (t *StructTyp) Add(name string, sym Sym) error {
	if sym.Type()!=SType{
		return ESymbolTypeError
	}
	t.members = append(t.members, NameTypPair{Name: name, Typ: sym.(TypSym)})
	return nil
}

func (t *StructTyp) SetParent(s SymTab) {
	t.parent = s
}

func (t *StructTyp) GetFieldIndex(name string) int {
	for i, m := range t.members {
		if m.Name==name{
			return i
		}
	}
	return -1
}

func (t *StructTyp) Type() symType {
	return SType
}

func (t *StructTyp) LlTyp() lltyp.Type {
	if t.llTyp==nil{
		var llTypes []lltyp.Type
		for _, m := range t.members {
			t:=m.Typ.LlTyp()
			if t==nil{
				return nil
			}
			llTypes = append(llTypes, t)
		}
		t.llTyp = lltyp.NewStruct(llTypes...)
	}
	return t.llTyp
}

type ValSym interface {
	Sym
	LlVal() llval.Value
	SetLlVal(llval.Value) error
	Typ() TypSym
}

type PrimaryVal struct {
	llVal llval.Value
	typ *PrimaryTyp
}

func (p *PrimaryVal) SetLlVal(value llval.Value) error {
	p.llVal=value
	return nil
}

func (p *PrimaryVal) Type() symType {
	return SValue
}

func (p *PrimaryVal) LlVal() llval.Value {
	return p.llVal
}

func (p *PrimaryVal) Typ() TypSym {
	return p.typ
}

type TupleVal struct {
	llVal   llval.Value
	lazyLlVal func()llval.Value
	Members []ValSym
	typ     *TupleTyp
	ParentSymTab
}

func (t *TupleVal) SetLlVal(value llval.Value) error {
	if t.llVal!=nil{
		return fmt.Errorf("assign to a tuple: %w", ENotSupport)
	}
	t.llVal = value
	return nil
}

func (t *TupleVal) SetLlValFunc(valueFunc func()llval.Value) {
	t.lazyLlVal=valueFunc
	return
}

func (t *TupleVal) Reduce() ValSym {
	switch len(t.Members) {
	case 0:
		return UnitVal
	case 1:
		return t.Members[0]
	default:
		mem:=make([]TypSym,0, len(t.Members))
		for _, s := range t.Members {
			mem = append(mem, s.Typ())
		}
		t.typ = GetTupleTyp(mem).(*TupleTyp)
		return t
	}
}

func (t *TupleVal) Type() symType {
	return SValue
}

func (t *TupleVal) LlVal() llval.Value {
	if t.llVal!=nil{
		return t.llVal
	}
	t.llVal=t.lazyLlVal()
	return t.llVal
}

func (t *TupleVal) Typ() TypSym {
	return t.typ
}

var tupleIdxReg *regexp.Regexp

func init() {
	tupleIdxReg=regexp.MustCompile("_\\d+")
}

func (t *TupleVal) GetField(name string, block *ir.Block) ValSym{
	if tupleIdxReg.MatchString(name){
		idx, err := strconv.Atoi(name[1:])
		if err != nil {
			return nil
		}
		if idx != -1 {
			f:=t.Members[idx]
			if f.LlVal()==nil{
				err := f.SetLlVal(block.NewExtractValue(t.llVal, uint64(idx)))
				if err != nil {
					log.Println(err)
					return nil
				}
			}
			return f
		}
	}
	return nil
}

func GetTupleTyp(types []TypSym)TypSym {
	switch len(types) {
	case 0:
		return UnitTyp
	case 1:
		return types[0]
	}
	if len(types) >10 {
		panic(ETupleTooLong)
	}
	var t [10]TypSym
	for i,it:=range types{
		t[i] = it
	}
	if s,ok:=tuplesPool[t];ok{
		return s
	}
	tp:=&TupleTyp{members: types}
	tuplesPool[t]=tp
	return tp
}

var tuplesPool = map[[10]TypSym]*TupleTyp{}


type StructVal struct {
	llVal llval.Value
	members []ValSym
	typ *StructTyp
	ParentSymTab
}

func (s *StructVal) SetLlVal(value llval.Value) error {
	return ENotSupport
}

func (s *StructVal) Type() symType {
	return SValue
}

func (s *StructVal) LlVal() llval.Value {
	return s.llVal
}

func (s *StructVal) Typ() TypSym {
	return s.typ
}

func (s *StructVal) Resolve(name string) Sym {
	idx:=s.typ.GetFieldIndex(name)
	if idx != -1 {
		return s.members[idx]
	}
	return s.parent.Resolve(name)
}

func (s *StructVal) GetField(name string, block *ir.Block) ValSym {
	idx:=s.typ.GetFieldIndex(name)
	if idx != -1 {
		f:=s.members[idx]
		if f.LlVal()==nil{
			err := f.SetLlVal(block.NewExtractValue(s.llVal, uint64(idx)))
			if err != nil {
				panic(err)
			}
		}
		return f
	}
	return nil
}


var UnitVal = &PrimaryVal{llVal: nil, typ: UnitTyp}
var TrueVal = &PrimaryVal{llVal: constant.True, typ: BoolTyp}
var FalseVal = &PrimaryVal{llVal: constant.False, typ: BoolTyp}
var UnitPhonyName = "%unit"