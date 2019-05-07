/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 19:34
 */
package conf

var Msg  = map[int]string{

	//default
	400001000: "表单参数解析失败,请检查代码",
	400001001: "控制器参数断言失败,请检查后再试",
	400001002: "数据类型转换失败,请检查后再试",
	400001003: "表单参数获取失败,请检查后再试",

	//post
	401000000: "文章标题不能为空,请检查后再试",
	401000001: "文章分类不能为空,请检查后再试",
	401000002: "文章标签不能为空,请检查后再试",
	401000003: "文章摘要不能为空,请检查后再试",
	401000004: "表单参数获取失败,请检查后再试",
	401000005: "图片上传失败,请检查后再试",

	//cate
	402000000: "查询分类数据失败,请检查后再试",
	402000001: "获取分类数据列表失败,请检查后再试",
	402000002: "分类名不能为空,请检查后再试",
	402000003: "分类别名不能为空,请检查后再试",
	402000004: "分类父级不能为空,请检查后再试",
	402000005: "分类描述不能为空,请检查后再试",
	402000006: "分类名长度超出范围,请检查后再试",
	402000007: "分类别名长度超出范围,请检查后再试",
	402000008: "分类描述长度超出范围,请检查后再试",
	402000009: "分类修改失败,请检查后再试",
	402000010: "分类创建失败,请检查后再试",
	402000011: "分类还有子节点,请检查后再试",

	//tag
	403000000: "标签名不能为空,请检查后再试",
	403000001: "标签名超过最大长度,请检查后再试",
	403000002: "标签别名不能为空,请检查后再试",
	403000003: "标签别名超过最大长度,请检查后再试",
	403000004: "标签描述不能为空,请检查后再试",
	403000005: "标签描述超过最大长度,请检查后再试",
	403000006: "创建标签失败,请检查后再试",
	403000007: "标签修改失败,请检查后再试",
	403000008: "标签查询失败,请检查后再试",

	//system
	405000000: "系统信息更新失败,请检查后再试",
	405000001: "请输入网站title",
	405000002: "网站Title长度超出最大长度,请检查后再试",
	405000003: "网站关键词不能为空,请检查后再试",
	405000004: "网站关键词长度超出最大长度,请检查后再试",
	405000005: "网站描述不能为空,请检查后再试",
	405000006: "网站描述长度超出最大范围,请检查后再试",
	405000007: "网站备案号不能为空,请检查后再试",
	405000008: "网站备案号长度超出,请检查后再试",
	405000009: "请选择网站前台主题,请检查后再试",
}

