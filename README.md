# CheckUp

在`Sandwich 1.3`中，健康检查的逻辑被提出单独作为一个checkup库。checkup库被sandwich-probe(>=1.3版本)和的Sandwich库(>=1.3版本)引用。
checkup处理内部消息中的`internal_command.CmdCheckup`，
也就是probe探针消息。Sandwich框架中的消息会在Server层的Checkup middleware时，由checkup进行字节数组的对比。
如果是checkup消息，就会被checkup库处理，调用逻辑层设置的`健康检查函数`，并向上返回。如果不是，就会继续往下传递。



