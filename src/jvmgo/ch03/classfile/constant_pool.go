package classfile

/**
3.常量池管理类
*/
type ConstantPool []ConstantInfo

//读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { //注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfp:
			i++ //占两个位置
		}
	}
	return cp
}

//按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cp != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

//从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getBameAndType(index uint16) (string, string) {
	nInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(nInfo.nameIndex)
	_type := self.getUtf8(nInfo.descriptionIndex)
	return name, _type
}

//从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(ConstantUtf8Info)
	return utf8Info.str
}
