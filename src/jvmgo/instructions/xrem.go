package instructions

import (
    "math"
    "jvmgo/rtda"
)

// Remainder double
type drem struct {NoOperandsInstruction}
func (self *drem) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := math.Mod(v1, v2) // todo
    stack.PushDouble(result)
}

// Remainder float
type frem struct {NoOperandsInstruction}
func (self *frem) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := float32(math.Mod(float64(v1), float64(v2))) // todo
    stack.PushFloat(result)
}

// Remainder int
type irem struct {NoOperandsInstruction}
func (self *irem) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 % v2
    stack.PushInt(result)
}

// Remainder long
type lrem struct {NoOperandsInstruction}
func (self *lrem) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 % v2
    stack.PushLong(result)
}
