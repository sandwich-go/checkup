# internalCmd

内部消息是与业务逻辑无关的消息集合。internalCmd是为了方便消息的添加、处理和测试，并能方便地提供给使用者。

内部消息使用checkup.InternalCmd作为消息结构定义，被包装的具体协议存放在Raw中。
````golang
type InternalCmd struct {
	Uri         string `json:"uri"`
	Raw         []byte `json:"raw"`
	PassThrough string `json:"passThrough,omitempty"`
}
````
内部消息为了在多种场合下使用，选择使用json进行编码。InternalCmd 和 Raw 都以json来序列化。
发送者只要能保证发送的格式符合 InternalCmd 和 Raw 以json格式解析即可

### 内部消息发送
内部消息的传输数据是一段json。是将checkup.InternalCmd Marshal成json。没有另外增加magic head，json开头为固定的{，可作为magic head做快速识别。
以probe发送CmdCheckup为例：
````golang
// 获取CmdCheckup的[]byte
b := checkup.NewInternalCmd(&internal_command.CmdCheckup{}, "123")
// 以TCP或http方式发送到目标服
ss.WriteFrame(b)
````
shell脚本发送：
````shell
curl -X POST -H  "acceptlication/json" -H "Content-Type: application/json" "http://127.0.0.1:8088" -d '{"uri":"internal_command.CmdCheckup","raw":"e30=","passThrough":"123"}' --connect-timeout 5
````

### 内部消息的接收处理
接收内部消息的服务，将接收到的未解析的字节流交给internalCmd处理，如果有回包消息，就是内部消息被处理了，将回包返回即可。
一个服务想使用已有的内部消息，也只需添加以下内容：
````golang
import "github.com/sandwich-go/checkup"

if resp, _ := checkup.HandleInternalCmd(context.Background(), bytesIn); resp != nil {
    return resp
}
````
具备内部消息特征的字节流段会被识别并处理，使用者不需要关心具体内容。

### 内部消息的添加
新增Internal Cmd时需要在 protocol/protos/internal_command/internal.proto 中添加内部消息的发包和回包定义，
发送符合包定义的json字符串。