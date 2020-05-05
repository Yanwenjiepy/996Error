package error_996

import "fmt"

const (
	Error996Str = "[996 is killing me]"
)

type Error996 struct {
	error996Content string
}

func Default() Error996 {

	return Error996{error996Content: Error996Str}
}

func New(error996Str string) Error996 {

	return Error996{error996Content: error996Str}
}

/*
	Wrap 包装错误并将包装后的错误返回
	description 为用于描述错误的string，应该只包含 %w 用于包装错误
	err 为待包装错误

	如果description中没有%w或者除了%w还有其他格式化参数则都将返回不可用的包装结果

	example:

	// 自定义996error消息头
	e := error_996.New("[Resist 996]")

	func decompressFile(fileName string) error {

		err := decompress(fileName)
		// 调用函数返回待包装错误
		if err != nil {
			// 构造description string,只包含%w,其他格式化参数提前处理
			descStr := fmt.Sprintf("decompress %v: ", fileName) + "%w"
			err = e.Wrap(descStr, err)
		}

		return err
	}


*/
func (e *Error996) Wrap(description string, err error) error {

	ef := e.error996Content + description
	err = fmt.Errorf(ef, err)

	return err
}
