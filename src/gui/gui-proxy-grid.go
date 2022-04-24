package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/golcl/lcl/types/messages"
	"strconv"
)

func (m *TGUIForm) proxyGrid() {
	//代理
	m.proxyLogsGrid = lcl.NewStringGrid(m)
	m.proxyLogsGrid.SetParent(m)
	m.proxyLogsGrid.SetFixedCols(0)
	m.proxyLogsGrid.SetBounds(0, m.height, m.width, 400)
	// 表格边框样式，这里设置为没有边框
	m.proxyLogsGrid.SetBorderStyle(types.BsNone)
	// 设置表格为平面样式
	m.proxyLogsGrid.SetFlat(true)
	// 表格列宽，自动大小后填充区域
	//m.proxyLogsGrid.SetAutoFillColumns(true)
	// 这里设置不可操作的列和行数
	m.proxyLogsGrid.SetFixedCols(0)
	//m.proxyLogsGrid.SetFixedRows(0)
	// 设置一些选项
	m.proxyLogsGrid.SetOptions(m.proxyLogsGrid.Options().Include(types.GoAlwaysShowEditor, types.GoCellHints, types.GoEditing, types.GoTabs))
	m.proxyLogsGrid.SetShowHint(true)
	// 设置不可操作列的背景颜色
	m.proxyLogsGrid.SetFixedColor(colors.ClGreen)

	// 设置flat后可以用这个修改fixed区域的表格线
	//m.proxyLogsGrid.SetFixedGridLineColor(colors.ClRed)
	//m.proxyLogsGrid.SetAnchors(types.NewSet(types.AkBottom, types.AkRight))
	m.proxyLogsGrid.SetVisible(true)

	var colNo = m.proxyLogsGrid.Columns().Add()
	colNo.SetWidth(100)
	colNo.Title().SetCaption("序号")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)
	colNo.SetReadOnly(true)
	//col1.SetReadOnly(true)

	var colAddr = m.proxyLogsGrid.Columns().Add()
	colAddr.SetWidth(m.width - 180)
	colAddr.Title().SetCaption("地址")
	colAddr.Title().SetAlignment(types.TaCenter)
	colAddr.SetReadOnly(true)

	var colDetailBtn = m.proxyLogsGrid.Columns().Add()
	colDetailBtn.SetWidth(60)
	colDetailBtn.SetButtonStyle(types.CbsButtonColumn)
	colDetailBtn.Title().SetCaption("详情")
	colDetailBtn.Title().SetAlignment(types.TaCenter)
	colDetailBtn.SetAlignment(types.TaCenter)

	m.proxyLogsGrid.SetOnButtonClick(m.onButtonClick)
	//m.proxyLogsGrid.SetOnSetEditText(m.onGridSetEditText)
	//选中列表某行列数据
	m.proxyLogsGrid.SetOnSelectCell(func(sender lcl.IObject, aCol, aRow int32, canSelect *bool) {
		selectRowIndex = aRow
		if aCol == 1 {
			selectCol2Value = m.proxyLogsGrid.Cells(aCol, aRow)
		}
	})
	//设置初始行数 1
	m.proxyLogsGrid.SetRowCount(rows)

	// 列表右键
	var pm = lcl.NewPopupMenu(m.proxyLogsGrid)
	var item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("复制")
	item.SetShortCutFromString("Ctrl+C")
	item.SetOnClick(func(lcl.IObject) {
		if selectCol2Value != "" {
			lcl.Clipboard.SetAsText(selectCol2Value)
		}
	})
	pm.Items().Add(item)
	item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("复制详情")
	item.SetShortCutFromString("Ctrl+Shift+C")
	item.SetOnClick(func(lcl.IObject) {
		if selectRowIndex != -1 {
			if row, ok := m.ProxyDetail[int(selectRowIndex)]; ok {
				if d, err := json.Marshal(row); err == nil {
					lcl.Clipboard.SetAsText(string(d))
				}
			}
		}
	})
	pm.Items().Add(item)
	item = lcl.NewMenuItem(m.proxyLogsGrid)
	item.SetCaption("清空")
	//item.SetShortCutFromString("")
	item.SetOnClick(func(lcl.IObject) {
		//m.proxyLogsGrid.Clear()
	})
	//pm.Items().Add(item)
	//添加到右键菜单
	m.proxyLogsGrid.SetPopupMenu(pm)

	//m.SetOnResize(func(sender lcl.IObject) {
	//	col2.SetWidth(m.Width() - 225)
	//})
	//m.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
	//	col2.SetWidth(m.Width() - 225)
	//})
}

//查看详情点击
func (m *TGUIForm) onButtonClick(sender lcl.IObject, aCol, aRow int32) {
	fmt.Println("click col", aCol, " row", aRow)
	if aCol == 2 { //查看详情点击
		if row, ok := m.ProxyDetail[int(aRow)]; ok {
			d, _ := json.Marshal(m.ProxyDetail[int(aRow)])
			fmt.Println("proxyDetail:", row.TargetUrl, " JSON:", string(d))
		}
	}
}

//func (m *TGUIForm) onGridSetEditText(sender lcl.IObject, col int32, row int32, value string) {
//	fmt.Println("onGridSetEditText", col, row, value)
//}

var (
	selectCol2Value string
	selectRowIndex  int32 = -1
	rows            int32 = 1
)

func (m *TGUIForm) AddProxyLogsGrid(proxyDetail *ProxyDetail) {
	lcl.ThreadSync(func() {
		rows++
		m.proxyLogsGrid.SetRowCount(rows)
		m.proxyLogsGrid.SetCells(0, rows-1, strconv.Itoa(proxyDetail.ID))
		m.proxyLogsGrid.SetCells(1, rows-1, proxyDetail.TargetUrl)
		m.proxyLogsGrid.SetCells(2, rows-1, "查看")
		m.proxyLogsGrid.Perform(messages.EM_SCROLLCARET, 7, 0)
	})
}

func (m *TGUIForm) SetProxyDetail(proxyDetail *ProxyDetail) {
	if _, ok := m.ProxyDetail[proxyDetail.ID]; !ok {
		m.AddProxyLogsGrid(proxyDetail)
	}
	m.ProxyDetail[proxyDetail.ID] = proxyDetail
}
