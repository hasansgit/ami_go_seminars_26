// Этот файл менять не нужно: тесты рассчитаны именно на эти объявления.

package treetransform

import "errors"

// ErrNilTransform сообщает, что Map вызвана без функции преобразования.
var ErrNilTransform = errors.New("nil transform")

// Transform преобразует значение одного узла.
type Transform func(int) (int, error)

// Node является узлом бинарного дерева.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}
