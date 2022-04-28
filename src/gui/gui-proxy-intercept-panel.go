package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/common"
	"gitee.com/snxamdf/http-server/src/entity"
	"path/filepath"
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
	TPanel         *lcl.TPanel
	ParamsGrid     *lcl.TStringGrid
	ParamsGridRow  int32
	HeadersGrid    *lcl.TStringGrid
	HeadersGridRow int32
	TBodyPanel     *ProxyInterceptRequestBodyPanel
}

//代理拦截请求Body Panel
type ProxyInterceptRequestBodyPanel struct {
	TPanel               *lcl.TPanel
	RawPanel             *lcl.TPanel
	RawMemo              *lcl.TMemo
	FormDataGridPanel    *lcl.TPanel
	FormDataGrid         *lcl.TStringGrid
	FormDataGridOpenFile *lcl.TOpenDialog
	FormDataGridList     map[int32]*entity.FormDataGridList
	FormDataGridRow      int32
}

//代理拦截响应Panel
type ProxyInterceptResponsePanel struct {
	TPanel *lcl.TPanel
}

//代理拦截配置Panel
type ProxyInterceptSettingPanel struct {
	TPanel *lcl.TPanel
}

//proxy intercept request UI
func (m *ProxyInterceptRequestPanel) initUI() {
	//Tabs 的控制标签
	resetPVars()
	pLeft = 0
	pTop = 0
	pHeight = m.TPanel.Height()/2 - 50
	pWidth = m.TPanel.Width()
	reqPageControl := lcl.NewPageControl(m.TPanel)
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	reqPageControl.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))

	//--- begin --- Request Query Params
	var paramsSheet = lcl.NewTabSheet(reqPageControl) //标签页
	paramsSheet.SetPageControl(reqPageControl)
	paramsSheet.SetCaption("　Request Query Params　")
	paramsSheet.SetAlign(types.AlClient)
	paramsPanel := lcl.NewPanel(m.TPanel) // 标签页
	paramsPanel.SetParent(paramsSheet)
	paramsPanel.SetAlign(types.AlClient)
	//按钮
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
			if m.ParamsGridRow > 1 {
				m.ParamsGrid.DeleteRow(aRow)
				m.ParamsGridRow--
			}
		}
	})
	m.RequestQueryParamsGridHead() //请求拦截参数表格头
	m.ParamsGrid.SetRow(m.ParamsGridRow)
	m.RequestQueryParamsGridAdd("", "") //默认添加一条
	//--- end --- Request Query Params

	//--- begin --- Request Headers
	var headersSheet = lcl.NewTabSheet(reqPageControl) //标签页
	headersSheet.SetPageControl(reqPageControl)
	headersSheet.SetCaption("　Request Headers　")
	headersSheet.SetAlign(types.AlClient)
	headersPanel := lcl.NewPanel(m.TPanel) // 标签页
	headersPanel.SetParent(headersSheet)
	headersPanel.SetAlign(types.AlClient)
	//按钮
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
			if m.HeadersGridRow > 1 {
				m.HeadersGrid.DeleteRow(aRow)
				m.HeadersGridRow--
			}
		}
	})
	m.HeaderGridHead()
	m.HeadersGrid.SetRow(m.HeadersGridRow)
	m.HeaderGridAdd("", "")
	//--- end --- Request Headers

	//--- begin --- Request Body
	resetPVars()
	pLeft = 0
	pTop = reqPageControl.Height()
	pHeight = m.TPanel.Height()/2 + 50
	pWidth = m.TPanel.Width()
	reqPageControl = lcl.NewPageControl(m.TPanel)
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	reqPageControl.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	var bodySheet = lcl.NewTabSheet(reqPageControl) //标签页
	bodySheet.SetPageControl(reqPageControl)
	bodySheet.SetCaption("　Request Body　")
	bodySheet.SetAlign(types.AlClient)
	m.TBodyPanel.TPanel = lcl.NewPanel(m.TPanel) // 标签页
	m.TBodyPanel.TPanel.SetParent(bodySheet)
	m.TBodyPanel.TPanel.SetAlign(types.AlClient)
	resetPVars()
	pLeft = 30
	pTop = 5
	var rdoRaw = lcl.NewRadioButton(m.TBodyPanel.TPanel)
	rdoRaw.SetParent(m.TBodyPanel.TPanel)
	rdoRaw.SetCaption("raw json")
	rdoRaw.SetLeft(pLeft)
	rdoRaw.SetTop(pTop)
	rdoRaw.SetOnClick(func(sender lcl.IObject) {
		m.TBodyPanel.bodyRdoCheckClick(0)
	})
	var rdoFormData = lcl.NewRadioButton(m.TBodyPanel.TPanel)
	rdoFormData.SetParent(m.TBodyPanel.TPanel)
	rdoFormData.SetCaption("form-data x-www-form-urlencoded binary")
	rdoFormData.SetLeft(rdoRaw.Left() + 120)
	rdoFormData.SetTop(pTop)
	rdoFormData.SetOnClick(func(sender lcl.IObject) {
		m.TBodyPanel.bodyRdoCheckClick(1)
	})

	m.TBodyPanel.initUI()
	//--- end --- Request Body

	//最后初始body选中
	rdoRaw.SetChecked(true)
	m.TBodyPanel.bodyRdoCheckClick(0)
}

//代理拦截请求Body Panel UI
func (m *ProxyInterceptRequestBodyPanel) initUI() {
	//raw
	resetPVars()
	pLeft = 10
	pTop = 31
	m.RawPanel = lcl.NewPanel(m.TPanel)
	m.RawPanel.SetParent(m.TPanel)
	m.RawPanel.SetBounds(pLeft, pTop, m.TPanel.Width()-20, m.TPanel.Height()-41)
	m.RawPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	m.RawPanel.SetVisible(false)
	m.RawMemo = lcl.NewMemo(m.RawPanel)
	m.RawMemo.SetParent(m.RawPanel)
	m.RawMemo.SetAlign(types.AlClient)

	//form-data & x-www-form-urlencoded
	m.FormDataGridOpenFile = lcl.NewOpenDialog(m.TPanel)
	m.FormDataGridOpenFile.SetTitle("选择上传文件")
	m.FormDataGridPanel = lcl.NewPanel(m.TPanel)
	m.FormDataGridPanel.SetParent(m.TPanel)
	m.FormDataGridPanel.SetBounds(pLeft, pTop, m.TPanel.Width()-20, m.TPanel.Height()-41)
	m.FormDataGridPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	m.FormDataGridPanel.SetVisible(false)
	m.FormDataGrid = lcl.NewStringGrid(m.FormDataGridPanel)
	m.FormDataGrid.SetParent(m.FormDataGridPanel)
	m.FormDataGrid.SetFixedCols(0)
	m.FormDataGrid.SetFixedColor(colors.ClGreen)
	m.FormDataGrid.SetBorderStyle(types.BsNone)
	m.FormDataGrid.SetFlat(true)
	m.FormDataGrid.SetOptions(m.FormDataGrid.Options().Include(types.GoAlwaysShowEditor, types.GoEditing))
	m.FormDataGrid.SetAlign(types.AlClient)
	m.FormDataGrid.SetOnSetEditText(func(sender lcl.IObject, aCol, aRow int32, value string) {
		row, ok := m.FormDataGridList[aRow]
		if !ok {
			row = &entity.FormDataGridList{}
			m.FormDataGridList[aRow] = row
		}
		if aCol == 0 { //列 类型
			if value == "Text" {
				m.FormDataGrid.SetCells(3, aRow, "--")
				row.FileValue = ""
			} else if value == "File" {
				if row.FileValue == "" {
					m.FormDataGrid.SetCells(3, aRow, "选择文件")
				} else {
					_, fileName := filepath.Split(row.FileValue)
					m.FormDataGrid.SetCells(3, aRow, fileName)
				}
			} else if value != "Text" && value != "File" {
				value = "Text"
				m.FormDataGrid.SetCells(0, aRow, value)
			}
			row.Type = value
		} else if aCol == 1 {
			row.Key = value
		} else if aCol == 2 {
			row.TextValue = value
		} else if aCol == 3 {
			//row.FileValue = value
		}
		if aCol == 1 || aCol == 2 {
			if aRow == m.FormDataGridRow-1 && (row.Key != "" || row.TextValue != "") {
				m.FormDataGridAdd("", "")
			}
		}
	})
	m.FormDataGrid.SetOnButtonClick(func(sender lcl.IObject, aCol, aRow int32) {
		fmt.Println("SetOnButtonClick", aRow)
		//按钮触发
		if aCol == 4 { //删除行
			if m.FormDataGridRow > 1 {
				m.FormDataGrid.DeleteRow(aRow)
				if _, ok := m.FormDataGridList[aRow]; ok {
					delete(m.FormDataGridList, aRow)
				}
				m.FormDataGridRow--
			}
		} else if aCol == 3 { //选择文件
			row, ok := m.FormDataGridList[aRow]
			if !ok {
				row = &entity.FormDataGridList{}
				m.FormDataGridList[aRow] = row
			}
			if row.Type == "File" {
				//解决同步到列表问题
				m.FormDataGrid.SetOptions(m.FormDataGrid.Options().Exclude(types.GoEditing))
				//解决同步到列表问题
				common.NewDebounce(1).Start(func() { //是个线程操作
					lcl.ThreadSync(func() { //需要主线程同步
						if m.FormDataGridOpenFile.Execute() {
							var filePath = m.FormDataGridOpenFile.FileName()
							row.FileValue = filePath
							_, fileName := filepath.Split(filePath)
							m.FormDataGrid.SetCells(3, aRow, fileName)
						}
					})
				})
				m.FormDataGrid.SetOptions(m.FormDataGrid.Options().Include(types.GoEditing))
			}
		}
	})
	m.FormDataGridHead()
	m.FormDataGrid.SetRow(m.FormDataGridRow)
	m.FormDataGridAdd("", "")
	//按钮
	var reqFormAddBtn = lcl.NewButton(m.FormDataGridPanel)
	reqFormAddBtn.SetParent(m.FormDataGridPanel)
	reqFormAddBtn.SetCaption("　添加参数　")
	reqFormAddBtn.SetBounds(520, 2, 80, 30)
	reqFormAddBtn.SetOnClick(func(sender lcl.IObject) {
		m.FormDataGridAdd("", "")
	})
}

//请求Body表格添加
func (m *ProxyInterceptRequestBodyPanel) FormDataGridAdd(key, value string) {
	lcl.ThreadSync(func() {
		m.FormDataGrid.InsertColRow(false, m.FormDataGridRow)
		m.FormDataGrid.SetCells(0, m.FormDataGridRow, "Text")
		m.FormDataGrid.SetCells(1, m.FormDataGridRow, key)
		m.FormDataGrid.SetCells(2, m.FormDataGridRow, value)
		m.FormDataGrid.SetCells(3, m.FormDataGridRow, "--")
		m.FormDataGrid.SetCells(4, m.FormDataGridRow, "删除")
		m.FormDataGridRow++
		m.FormDataGrid.SetRowCount(m.FormDataGridRow)
	})
}

//请求Body表格头
func (m *ProxyInterceptRequestBodyPanel) FormDataGridHead() {
	var colType = m.FormDataGrid.Columns().Add()
	colType.SetWidth(50)
	colType.Title().SetCaption("TYPE")
	colType.SetButtonStyle(types.CbsPickList)
	colType.Title().SetAlignment(types.TaCenter)
	colType.SetAlignment(types.TaCenter)
	colType.PickList().Add("Text")
	colType.PickList().Add("File")

	var colKey = m.FormDataGrid.Columns().Add()
	colKey.SetWidth(150)
	colKey.Title().SetCaption("Key")

	var colTextValue = m.FormDataGrid.Columns().Add()
	colTextValue.SetWidth(150)
	colTextValue.Title().SetCaption("Text Value")

	var colFileValue = m.FormDataGrid.Columns().Add()
	colFileValue.SetWidth(100)
	colFileValue.Title().SetCaption("Select File")
	colFileValue.Title().SetAlignment(types.TaCenter)
	colFileValue.SetButtonStyle(types.CbsButtonColumn)
	colFileValue.SetAlignment(types.TaCenter)

	var delBtn = m.FormDataGrid.Columns().Add()
	delBtn.SetWidth(60)
	delBtn.Title().SetCaption("操作")
	delBtn.Title().SetAlignment(types.TaCenter)
	delBtn.SetButtonStyle(types.CbsButtonColumn)
	delBtn.SetAlignment(types.TaCenter)
}

//body radio 按钮点击
func (m *ProxyInterceptRequestBodyPanel) bodyRdoCheckClick(t int) {
	fmt.Println(t)
	m.RawPanel.SetVisible(t == 0)
	m.FormDataGridPanel.SetVisible(t == 1)
}

//请求拦截头添加
func (m *ProxyInterceptRequestPanel) HeaderGridAdd(key, value string) {
	lcl.ThreadSync(func() {
		m.HeadersGrid.InsertColRow(false, m.HeadersGridRow)
		m.HeadersGrid.SetCells(0, m.HeadersGridRow, "1")
		m.HeadersGrid.SetCells(1, m.HeadersGridRow, key)
		m.HeadersGrid.SetCells(2, m.HeadersGridRow, value)
		m.HeadersGrid.SetCells(3, m.HeadersGridRow, "删除")
		m.HeadersGridRow++
		m.HeadersGrid.SetRowCount(m.HeadersGridRow)
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
	colNo.Title().SetCaption("Key")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)

	var colAddr = m.HeadersGrid.Columns().Add()
	colAddr.SetWidth(180)
	colAddr.Title().SetCaption("Value")
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
		m.ParamsGrid.InsertColRow(false, m.ParamsGridRow)
		m.ParamsGrid.SetCells(0, m.ParamsGridRow, "1")
		m.ParamsGrid.SetCells(1, m.ParamsGridRow, key)
		m.ParamsGrid.SetCells(2, m.ParamsGridRow, value)
		m.ParamsGrid.SetCells(3, m.ParamsGridRow, "删除")
		m.ParamsGridRow++
		m.ParamsGrid.SetRowCount(m.ParamsGridRow)
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
	colNo.Title().SetCaption("Key")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)

	var colAddr = m.ParamsGrid.Columns().Add()
	colAddr.SetWidth(180)
	colAddr.Title().SetCaption("Value")
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
