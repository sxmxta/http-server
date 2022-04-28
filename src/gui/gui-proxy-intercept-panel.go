package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
)

//代理拦截配置Panel
type ProxyInterceptPanel struct {
	TPanel                      *lcl.TPanel
	StateLabel                  *lcl.TLabel                  //拦截状态
	ProxyInterceptRequestPanel  *ProxyInterceptRequestPanel  //代理拦截请求Panel
	ProxyInterceptResponsePanel *ProxyInterceptResponsePanel //代理拦截响应Panel
	ProxyInterceptSettingPanel  *ProxyInterceptSettingPanel  //代理拦截配置Panel
}

//代理拦截请求Panel
type ProxyInterceptRequestPanel struct {
	TPanel      *lcl.TPanel
	ParamsGrid  *lcl.TStringGrid
	ParamsRow   int32
	HeadersGrid *lcl.TStringGrid
	HeadersRow  int32
	TBodyPanel  *lcl.TPanel
}

//代理拦截响应Panel
type ProxyInterceptResponsePanel struct {
	TPanel *lcl.TPanel
}

//代理拦截配置Panel
type ProxyInterceptSettingPanel struct {
	TPanel *lcl.TPanel
}

//request
func (m *ProxyInterceptRequestPanel) initUI() {
	resetPVars()
	pLeft = 0
	pTop = 0
	pWidth = m.TPanel.Width()
	pHeight = m.TPanel.Height()

	//Tabs 的控制标签
	reqPageControl := lcl.NewPageControl(m.TPanel)
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	reqPageControl.SetAlign(types.AlClient)

	//--- begin --- Request Query Params
	var paramsSheet = lcl.NewTabSheet(reqPageControl) //标签页
	paramsSheet.SetPageControl(reqPageControl)
	paramsSheet.SetCaption("　Request Query Params　")
	paramsSheet.SetAlign(types.AlClient)
	paramsPanel := lcl.NewPanel(m.TPanel) // 标签页
	paramsPanel.SetParent(paramsSheet)
	paramsPanel.SetBounds(0, 0, pWidth, pHeight)
	paramsPanel.SetAlign(types.AlClient)
	var reqQueryParamAddBtn = lcl.NewButton(m.TPanel)
	reqQueryParamAddBtn.SetParent(paramsSheet)
	reqQueryParamAddBtn.SetCaption("　添加参数　")
	reqQueryParamAddBtn.SetBounds(460, 1, 60, 30)
	reqQueryParamAddBtn.SetOnClick(func(sender lcl.IObject) {
		m.RequestQueryParamsGridAdd("", "")
	})

	//ParamsGrid
	m.ParamsGrid = lcl.NewStringGrid(paramsPanel)
	m.ParamsGrid.SetParent(paramsPanel)
	m.ParamsGrid.SetFixedCols(0)
	m.ParamsGrid.SetFixedColor(colors.ClGreen)
	m.ParamsGrid.SetAlign(types.AlClient)
	m.ParamsGrid.SetBorderStyle(types.BsNone)
	m.ParamsGrid.SetFlat(true)
	m.ParamsGrid.SetOptions(m.ParamsGrid.Options().Include(types.GoAlwaysShowEditor, types.GoCellHints, types.GoEditing, types.GoTabs))
	m.ParamsGrid.SetOnButtonClick(func(sender lcl.IObject, aCol, aRow int32) {
		if aCol == 3 {
			if m.ParamsRow > 1 {
				m.ParamsGrid.DeleteRow(aRow)
				m.ParamsRow--
			}
		}
	})
	m.RequestQueryParamsGridHead() //请求拦截参数表格头
	m.ParamsGrid.SetRow(m.ParamsRow)
	m.RequestQueryParamsGridAdd("", "") //默认添加一条
	//--- end --- Request Query Params

	//--- begin --- Request Headers
	var headersSheet = lcl.NewTabSheet(reqPageControl) //标签页
	headersSheet.SetPageControl(reqPageControl)
	headersSheet.SetCaption("　Request Headers　")
	headersSheet.SetAlign(types.AlClient)
	headersPanel := lcl.NewPanel(m.TPanel) // 标签页
	headersPanel.SetParent(headersSheet)
	headersPanel.SetBounds(0, 0, pWidth, pHeight)
	headersPanel.SetAlign(types.AlClient)
	var reqHeaderAddBtn = lcl.NewButton(m.TPanel)
	reqHeaderAddBtn.SetParent(headersSheet)
	reqHeaderAddBtn.SetCaption("　添加请求头　")
	reqHeaderAddBtn.SetBounds(460, 1, 80, 30)
	reqHeaderAddBtn.SetOnClick(func(sender lcl.IObject) {
		m.HeaderGridAdd("", "")
	})
	//HeadersGrid
	m.HeadersGrid = lcl.NewStringGrid(headersPanel)
	m.HeadersGrid.SetParent(headersPanel)
	m.HeadersGrid.SetFixedCols(0)
	m.HeadersGrid.SetFixedColor(colors.ClGreen)
	m.HeadersGrid.SetAlign(types.AlClient)
	m.HeadersGrid.SetBorderStyle(types.BsNone)
	m.HeadersGrid.SetFlat(true)
	m.HeadersGrid.SetOptions(m.HeadersGrid.Options().Include(types.GoAlwaysShowEditor, types.GoCellHints, types.GoEditing, types.GoTabs))
	m.HeadersGrid.SetOnButtonClick(func(sender lcl.IObject, aCol, aRow int32) {
		if aCol == 3 {
			if m.HeadersRow > 1 {
				m.HeadersGrid.DeleteRow(aRow)
				m.HeadersRow--
			}
		}
	})
	m.HeaderGridHead()
	m.HeadersGrid.SetRow(m.HeadersRow)
	m.HeaderGridAdd("", "")
	//--- end --- Request Headers

	//--- begin --- Request Body
	var bodySheet = lcl.NewTabSheet(reqPageControl) //标签页
	bodySheet.SetPageControl(reqPageControl)
	bodySheet.SetCaption("　Request Body　")
	bodySheet.SetAlign(types.AlClient)
	m.TBodyPanel = lcl.NewPanel(m.TPanel) // 标签页
	m.TBodyPanel.SetParent(bodySheet)
	m.TBodyPanel.SetBounds(0, 0, pWidth, pHeight)
	m.TBodyPanel.SetAlign(types.AlClient)
	resetPVars()
	pLeft = 30
	pTop = 5
	var rdoRaw = lcl.NewRadioButton(m.TBodyPanel)
	rdoRaw.SetParent(m.TBodyPanel)
	rdoRaw.SetCaption("raw")
	rdoRaw.SetLeft(pLeft)
	rdoRaw.SetTop(pTop)
	rdoRaw.SetOnClick(func(sender lcl.IObject) {
		m.bodyRdoCheckClick(0)
	})
	var rdoFormData = lcl.NewRadioButton(m.TBodyPanel)
	rdoFormData.SetParent(m.TBodyPanel)
	rdoFormData.SetCaption("form-data")
	rdoFormData.SetLeft(rdoRaw.Left() + 60)
	rdoFormData.SetTop(pTop)
	rdoFormData.SetOnClick(func(sender lcl.IObject) {
		m.bodyRdoCheckClick(1)
	})
	var rdoXWWWUrlEncode = lcl.NewRadioButton(m.TBodyPanel)
	rdoXWWWUrlEncode.SetParent(m.TBodyPanel)
	rdoXWWWUrlEncode.SetCaption("x-www-form-urlencoded")
	rdoXWWWUrlEncode.SetLeft(rdoFormData.Left() + 90)
	rdoXWWWUrlEncode.SetTop(pTop)
	rdoXWWWUrlEncode.SetOnClick(func(sender lcl.IObject) {
		m.bodyRdoCheckClick(2)
	})
	//--- end --- Request Body
}

//body radio 按钮点击
func (m *ProxyInterceptRequestPanel) bodyRdoCheckClick(t int) {
	fmt.Println(t)
}

//请求拦截头添加
func (m *ProxyInterceptRequestPanel) HeaderGridAdd(key, value string) {
	lcl.ThreadSync(func() {
		m.HeadersGrid.InsertColRow(false, m.HeadersRow)
		m.HeadersGrid.SetCells(0, m.HeadersRow, "1")
		m.HeadersGrid.SetCells(1, m.HeadersRow, key)
		m.HeadersGrid.SetCells(2, m.HeadersRow, value)
		m.HeadersGrid.SetCells(3, m.HeadersRow, "删除")
		m.HeadersRow++
		m.HeadersGrid.SetRowCount(m.HeadersRow)
	})
}

//请求拦截参数表格头
func (m *ProxyInterceptRequestPanel) HeaderGridHead() {
	var chkBox = m.HeadersGrid.Columns().Add()
	chkBox.SetWidth(30)
	chkBox.SetButtonStyle(types.CbsCheckboxColumn)
	chkBox.Title().SetCaption("启用")

	var colNo = m.HeadersGrid.Columns().Add()
	colNo.SetWidth(180)
	colNo.Title().SetCaption("KEY")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)

	var colAddr = m.HeadersGrid.Columns().Add()
	colAddr.SetWidth(180)
	colAddr.Title().SetCaption("VALUE")
	colAddr.Title().SetAlignment(types.TaCenter)
	colAddr.SetAlignment(types.TaCenter)

	var delBtn = m.HeadersGrid.Columns().Add()
	delBtn.SetWidth(60)
	delBtn.Title().SetCaption("操作")
	delBtn.Title().SetAlignment(types.TaCenter)
	delBtn.SetButtonStyle(types.CbsButtonColumn)
	delBtn.SetAlignment(types.TaCenter)
}

//请求拦截参数列表添加
func (m *ProxyInterceptRequestPanel) RequestQueryParamsGridAdd(key, value string) {
	lcl.ThreadSync(func() {
		m.ParamsGrid.InsertColRow(false, m.ParamsRow)
		m.ParamsGrid.SetCells(0, m.ParamsRow, "1")
		m.ParamsGrid.SetCells(1, m.ParamsRow, key)
		m.ParamsGrid.SetCells(2, m.ParamsRow, value)
		m.ParamsGrid.SetCells(3, m.ParamsRow, "删除")
		m.ParamsRow++
		m.ParamsGrid.SetRowCount(m.ParamsRow)
	})
}

//请求拦截参数表格头
func (m *ProxyInterceptRequestPanel) RequestQueryParamsGridHead() {
	var chkBox = m.ParamsGrid.Columns().Add()
	chkBox.SetWidth(30)
	chkBox.SetButtonStyle(types.CbsCheckboxColumn)
	chkBox.Title().SetCaption("启用")

	var colNo = m.ParamsGrid.Columns().Add()
	colNo.SetWidth(180)
	colNo.Title().SetCaption("KEY")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)

	var colAddr = m.ParamsGrid.Columns().Add()
	colAddr.SetWidth(180)
	colAddr.Title().SetCaption("VALUE")
	colAddr.Title().SetAlignment(types.TaCenter)
	colAddr.SetAlignment(types.TaCenter)

	var delBtn = m.ParamsGrid.Columns().Add()
	delBtn.SetWidth(60)
	delBtn.Title().SetCaption("操作")
	delBtn.Title().SetAlignment(types.TaCenter)
	delBtn.SetButtonStyle(types.CbsButtonColumn)
	delBtn.SetAlignment(types.TaCenter)
}

//response
func (m *ProxyInterceptResponsePanel) initUI() {
	resetPVars()
	pLeft = 0
	pTop = 25
	pWidth = m.TPanel.Width()
	pHeight = m.TPanel.Height()

	reqPageControl := lcl.NewPageControl(m.TPanel) //Tabs 的控制标签
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	reqPageControl.SetAlign(types.AlClient)

	sheet := lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Response Headers　")
	sheet.SetAlign(types.AlClient)
	headersPanel := lcl.NewPanel(m.TPanel) // 标签页
	headersPanel.SetParent(sheet)
	headersPanel.SetBounds(0, 0, pWidth, pHeight)
	headersPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Response Body　")
	sheet.SetAlign(types.AlClient)
	bodyPanel := lcl.NewPanel(m.TPanel) // 标签页
	bodyPanel.SetParent(sheet)
	bodyPanel.SetBounds(0, 0, pWidth, pHeight)
	bodyPanel.SetAlign(types.AlClient)
}

//setting
func (m *ProxyInterceptSettingPanel) initUI() {

}

//代理拦截配置Panel
func (m *ProxyInterceptPanel) initUI() {
	resetPVars()
	pLeft = 0
	pTop = 0
	pWidth = m.TPanel.Width()
	pHeight = m.TPanel.Height()

	reqPageControl := lcl.NewPageControl(m.TPanel) //Tabs 的控制标签
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	reqPageControl.SetAlign(types.AlClient)

	sheetInterReq := lcl.NewTabSheet(reqPageControl) //标签页
	sheetInterReq.SetPageControl(reqPageControl)
	sheetInterReq.SetCaption("　拦截请求　")
	sheetInterReq.SetAlign(types.AlClient)
	m.ProxyInterceptRequestPanel.TPanel = lcl.NewPanel(m.TPanel) //ProxyInterceptRequestPanel 标签页
	m.ProxyInterceptRequestPanel.TPanel.SetParent(sheetInterReq)
	m.ProxyInterceptRequestPanel.TPanel.SetBounds(0, 0, pWidth, pHeight)
	m.ProxyInterceptRequestPanel.TPanel.SetAlign(types.AlClient)

	sheetInterRes := lcl.NewTabSheet(reqPageControl) //标签页
	sheetInterRes.SetPageControl(reqPageControl)
	sheetInterRes.SetCaption("　拦截响应　")
	sheetInterRes.SetAlign(types.AlClient)
	m.ProxyInterceptResponsePanel.TPanel = lcl.NewPanel(m.TPanel) //responsePanel 标签页
	m.ProxyInterceptResponsePanel.TPanel.SetParent(sheetInterRes)
	m.ProxyInterceptResponsePanel.TPanel.SetBounds(0, 0, pWidth, pHeight)
	m.ProxyInterceptResponsePanel.TPanel.SetAlign(types.AlClient)

	sheetInterSet := lcl.NewTabSheet(reqPageControl) //标签页
	sheetInterSet.SetPageControl(reqPageControl)
	sheetInterSet.SetCaption("　拦截配置　")
	sheetInterSet.SetAlign(types.AlClient)
	m.ProxyInterceptSettingPanel.TPanel = lcl.NewPanel(m.TPanel) //responsePanel 标签页
	m.ProxyInterceptSettingPanel.TPanel.SetParent(sheetInterSet)
	m.ProxyInterceptSettingPanel.TPanel.SetBounds(0, 0, pWidth, pHeight)
	m.ProxyInterceptSettingPanel.TPanel.SetAlign(types.AlClient)

	//初始化子组件
	m.ProxyInterceptRequestPanel.initUI()
	m.ProxyInterceptResponsePanel.initUI()
	m.ProxyInterceptSettingPanel.initUI()

	//TODO 测试 tabs 切换
	testBtn := lcl.NewButton(m.TPanel)
	testBtn.SetParent(m.TPanel)
	testBtn.SetCaption("测试切换")
	testBtn.SetLeft(m.TPanel.Width() - 100)
	var is int32 = 1
	testBtn.SetOnClick(func(sender lcl.IObject) {
		reqPageControl.SetActivePageIndex(is)
		is++
		if is > 2 {
			is = 0
		}
	})
}
