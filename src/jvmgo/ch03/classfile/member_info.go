package classfile

/**
3.定义结构体，统一表示字段和方法
*/

type MemberInfo struct {
	cp               ConstantPool
	accessFlags      uint16
	nameIndex        uint16
	descriptionIndex uint16
	attributes       []AttributeInfo
}

//readMembers()读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

//readMember()函数读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:               cp,
		accessFlags:      reader.readUint16(),
		nameIndex:        reader.readUint16(),
		descriptionIndex: reader.readUint16(),
		attributes:       readAttributes(reader, cp),
	}
}
func (self *MemberInfo) AccessFlags() uint16 {}

//从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

//从常量池查找字段或方法描述符
func (self *MemberInfo) Description() string {
	return self.cp.getUtf8(self.descriptionIndex)
}
