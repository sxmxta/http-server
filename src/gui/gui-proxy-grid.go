package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/entity"
	"sync"
)

func (m *TGUIForm) proxyGrid() {
	//代理
	m.proxyLogsGrid = lcl.NewStringGrid(m)
	m.proxyLogsGrid.SetParent(m)
	m.proxyLogsGrid.SetScrollBars(types.SsAutoBoth)
	var mh = m.Height()
	mh = mh - uiHeight - 20
	m.proxyLogsGrid.SetBounds(0, uiHeight, uiWidth, mh)
	//m.proxyLogsGrid.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop))
	// 表格边框样式，这里设置为没有边框
	m.proxyLogsGrid.SetBorderStyle(types.BsNone)
	m.proxyLogsGrid.SetColor(colors.ClWhite)
	// 设置表格为平面样式
	m.proxyLogsGrid.SetFlat(true)
	// 表格列宽，自动大小后填充区域
	//m.proxyLogsGrid.SetAutoFillColumns(true)
	// 这里设置不可操作的列和行数
	m.proxyLogsGrid.SetFixedCols(0)
	//m.proxyLogsGrid.SetFixedRows(0)
	// 设置一些选项
	//m.proxyLogsGrid.SetOptions(m.proxyLogsGrid.Options().Include(types.GoRowHighlight))
	// 设置不可操作列的背景颜色
	m.proxyLogsGrid.SetFixedColor(colors.ClGreen)
	//var topRow int32 = 1
	m.proxyLogsGrid.SetOnMouseWheelUp(func(sender lcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		m.proxyLogsGrid.SetTopRow(1)
	})
	m.proxyLogsGrid.SetOnMouseWheelDown(func(sender lcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		m.proxyLogsGrid.SetTopRow(m.proxyLogsGridCountRow)
	})
	// 设置flat后可以用这个修改fixed区域的表格线
	//m.proxyLogsGrid.SetFixedGridLineColor(colors.ClRed)
	//m.proxyLogsGrid.SetAnchors(types.NewSet(types.AkBottom, types.AkRight))
	//m.proxyLogsGrid.SetVisible(false)
	//绘制表格头
	m.proxyLogsGridHead()
	var bgColor = m.proxyLogsGrid.Color()
	m.proxyLogsGrid.SetOnDrawCell(func(sender lcl.IObject, aCol, aRow int32, aRect types.TRect, state types.TGridDrawState) {
		if aRow < 1 {
			return
		}
		//fmt.Println( aCol, aRow)
		if aCol == 2 || aCol == 3 {
			var row = m.proxyLogsGridCountRow - aRow
			//fmt.Println("col", aCol, " aRow", aRow, "  总行", m.proxyLogsGridCountRow, "  实际", row)
			//row是反正来的
			if rowStyle, ok := m.ProxyLogsGridRowStyle[row]; ok {
				for col, style := range rowStyle.GetCols() {
					if col == aCol && style != nil {
						if style.IsColor() {
							m.proxyLogsGrid.Canvas().Brush().SetColor(bgColor)
							m.proxyLogsGrid.Canvas().FillRect(aRect)
							m.proxyLogsGrid.Canvas().Font().SetColor(style.TColor())
							m.proxyLogsGrid.Canvas().TextRect2(&aRect, style.Text(), types.NewSet(types.TfCenter))
							//fmt.Println(row, aCol, style.Text())
						}
					}
				}
			}
			//TODO 以下注释代码，参考使用。
			//m.proxyLogsGrid.Canvas().Brush().SetTColor(colors.ClRed)//设置刷子颜色
			//m.proxyLogsGrid.Canvas().FillRect(aRect)//将表格全部刷新
			//m.proxyLogsGrid.Canvas().Font().SetTColor(colors.ClRed)//设置字体颜色
			//m.proxyLogsGrid.Canvas().TextRect(aRect, aRect.Left, aRect.Top, "测试测试")// 1 画文字
			//m.proxyLogsGrid.Canvas().TextRect2(&aRect,"测试测试",types.NewSet(types.TfCenter))// 2 画文字 带有字体设置
			//m.proxyLogsGrid.Canvas().TextOut(aRect.Left,aRect.Top,"测试测试")// 3 画文字
		}
	})
	//鼠标事件 up
	m.proxyLogsGrid.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbRight { //鼠标右键
			var point = types.TPoint{}
			point.X = x
			point.Y = y
			//获取鼠标右键位置的表格，返回的值 x=列 y=行
			var cell = m.proxyLogsGrid.MouseToCell(point)
			if cell.Y < 1 {
				return
			}
			var value = m.proxyLogsGrid.Cells(cell.X, cell.Y) //取值
			fmt.Println(value)
		}
	})

	//选中列表某行列数据
	m.proxyLogsGrid.SetOnSelectCell(func(sender lcl.IObject, aCol, aRow int32, canSelect *bool) {
		//fmt.Println("SetOnSelectCell", aRow)
		if aRow > 0 {
			//实际的ProxyDetails map对应的key
			logGridSelDetailKey = m.proxyLogsGridCountRow - aRow
			//放到剪切版
			logGridSelCol2Value = m.proxyLogsGrid.Cells(aCol, aRow)
			m.selectGridUpdate()
		}
	})
	//设置初始行数 1
	m.proxyLogsGrid.SetRowCount(m.proxyLogsGridCountRow)

	// 列表右键
	var pm = lcl.NewPopupMenu(m.proxyLogsGrid)
	var item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("复制地址")
	item.SetShortCutFromString("Ctrl+C")
	item.SetOnClick(func(lcl.IObject) {
		if logGridSelCol2Value != "" {
			lcl.Clipboard.SetAsText(logGridSelCol2Value)
		}
	})
	pm.Items().Add(item)
	item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("复制详情")
	item.SetShortCutFromString("Ctrl+Shift+C")
	item.SetOnClick(func(lcl.IObject) {
		if logGridSelDetailKey != -1 {
			if row, ok := m.ProxyDetails[logGridSelDetailKey]; ok {
				if d, err := json.Marshal(row); err == nil {
					lcl.Clipboard.SetAsText(string(d))
				}
			}
		}
	})
	pm.Items().Add(item)
	item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("添加拦截")
	item.SetShortCutFromString("Ctrl+Shift+A")
	item.SetOnClick(func(lcl.IObject) {
		if logGridSelDetailKey != -1 && logGridSelDetailKey > 0 {
			if row, ok := m.ProxyDetails[logGridSelDetailKey]; ok {
				fmt.Println(row.TargetUrl)
				//m.ProxyDetailUI.ProxyInterceptConfigPanel.AddProxyInterceptConfig(value, true)
			}
		}
	})
	pm.Items().Add(item)
	//item = lcl.NewMenuItem(m.proxyLogsGrid)
	//item.SetCaption("清空")
	////item.SetShortCutFromString("")
	//item.SetOnClick(func(lcl.IObject) {
	//	m.proxyLogsGridClear()
	//})
	//pm.Items().Add(item)

	//添加到右键菜单
	m.proxyLogsGrid.SetPopupMenu(pm)
}

func (m *TGUIForm) selectGridUpdate() {
	if rowData, ok := m.ProxyDetails[logGridSelDetailKey]; ok {
		m.ProxyDetailUI.RequestDetailViewPanel.updateRequestSheet(rowData)
		m.ProxyDetailUI.RequestDetailViewPanel.updateResponseSheet(rowData)
		fmt.Println(rowData.Method)
	}
}

//代理日志grid表格头
func (m *TGUIForm) proxyLogsGridHead() {
	var colNo = m.proxyLogsGrid.Columns().Add()
	colNo.SetWidth(60)
	colNo.Title().SetCaption("序号")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)
	colNo.SetReadOnly(true)
	//col1.SetReadOnly(true)

	var colAddr = m.proxyLogsGrid.Columns().Add()
	colAddr.SetWidth(uiWidth - 200)
	colAddr.Title().SetCaption("地址 - (右键菜单复制)")
	colAddr.Title().SetAlignment(types.TaCenter)
	colAddr.SetReadOnly(true)

	var colState = m.proxyLogsGrid.Columns().Add()
	colState.SetWidth(60)
	colState.Title().SetCaption("流程")
	colState.Title().SetAlignment(types.TaCenter)
	colState.SetReadOnly(true)

	var colDetailBtn = m.proxyLogsGrid.Columns().Add()
	colDetailBtn.SetWidth(60)
	//colDetailBtn.SetButtonStyle(types.CbsButtonColumn)
	colDetailBtn.Title().SetCaption("状态码")
	colDetailBtn.Title().SetAlignment(types.TaCenter)
	colDetailBtn.SetAlignment(types.TaCenter)
}

var (
	logGridSelCol2Value string      //选中表格第二列的值
	logGridSelDetailKey int32  = -1 //选中表格行下标
	logGridInsertRow    int32  = 1  //在第指定行插入
)

//代理日志grid添加一行
func (m *TGUIForm) proxyLogsGridAdd(proxyDetail *entity.ProxyDetail) {
	m.proxyLogsGridCountRow = proxyDetail.ID + 1
	lcl.ThreadSync(func() {
		//proxyDetail.Row = m.proxyLogsGridCountRow
		//在指定行插入行
		m.proxyLogsGrid.InsertColRow(false, logGridInsertRow)
		//给指定行的列设置值
		m.proxyLogsGrid.SetCells(0, logGridInsertRow, fmt.Sprintf("%v", proxyDetail.ID))
		m.proxyLogsGrid.SetCells(1, logGridInsertRow, proxyDetail.Method+" - "+proxyDetail.TargetUrl)
		//text, _ := proxyDetail.GetState()
		m.proxyLogsGrid.SetCells(2, logGridInsertRow, "")
		m.proxyLogsGrid.SetCells(3, logGridInsertRow, fmt.Sprintf("%v", proxyDetail.StateCode))
		//给表格设置新总行数
		m.proxyLogsGrid.SetRowCount(m.proxyLogsGridCountRow)
	})
}

//代理日志grid清除-还未完善
func (m *TGUIForm) proxyLogsGridClear() {
	m.proxyLogsGrid.Clear()
	logGridSelCol2Value = ""    //选中列的值
	logGridSelDetailKey = -1    //实际的ProxyDetails map对应的key
	m.proxyLogsGridCountRow = 1 //总行数
	m.proxyLogsGrid.SetRowCount(m.proxyLogsGridCountRow)
	//m.proxyLogsGridHead()
}

//设置代理详情锁
var setProxyDetailLock = sync.RWMutex{}

//设置代理详情
func (m *TGUIForm) setProxyDetail(proxyDetail *entity.ProxyDetail) {
	setProxyDetailLock.Lock()
	defer setProxyDetailLock.Unlock()
	//添加到 代理详情数据集合
	if _, ok := m.ProxyDetails[proxyDetail.ID]; !ok {
		//代理日志grid添加一行
		m.proxyLogsGridAdd(proxyDetail)
	}
	//更新表格颜色参数设置 SetOnDrawCell 事件
	text, color := proxyDetail.GetState()
	if rowStyle, ok := m.ProxyLogsGridRowStyle[proxyDetail.ID]; ok {
		rowStyle.GetColStyle(2).SetTColor(color).SetText(text)
		rowStyle.GetColStyle(3).SetTColor(color).SetText(fmt.Sprintf("%v", proxyDetail.StateCode))
	} else {
		rowStyle = entity.NewRowStyle()
		rowStyle.GetColStyle(2).SetTColor(color).SetText(text)
		rowStyle.GetColStyle(3).SetTColor(color).SetText(fmt.Sprintf("%v", proxyDetail.StateCode))
		m.ProxyLogsGridRowStyle[proxyDetail.ID] = rowStyle
	}
	m.proxyLogsGrid.Invalidate()
	//m.proxyLogsGrid.SetCells(2, row, " ")
	//更新集合内容
	m.ProxyDetails[proxyDetail.ID] = proxyDetail
	m.selectGridUpdate() //更新右侧选中数据

}
