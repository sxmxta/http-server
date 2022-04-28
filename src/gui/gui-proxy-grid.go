package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/entity"
	"sync"
	"sync/atomic"
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
	// 设置表格为平面样式
	m.proxyLogsGrid.SetFlat(true)
	// 表格列宽，自动大小后填充区域
	//m.proxyLogsGrid.SetAutoFillColumns(true)
	// 这里设置不可操作的列和行数
	m.proxyLogsGrid.SetFixedCols(0)
	//m.proxyLogsGrid.SetFixedRows(0)
	// 设置一些选项
	m.proxyLogsGrid.SetOptions(m.proxyLogsGrid.Options().Include(types.GoAlwaysShowEditor, types.GoCellHints, types.GoEditing, types.GoTabs))
	// 设置不可操作列的背景颜色
	m.proxyLogsGrid.SetFixedColor(colors.ClGreen)
	//var topRow int32 = 1
	m.proxyLogsGrid.SetOnMouseWheelUp(func(sender lcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		m.proxyLogsGrid.SetTopRow(1)
	})
	m.proxyLogsGrid.SetOnMouseWheelDown(func(sender lcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		m.proxyLogsGrid.SetTopRow(m.proxyLogsGridRow)
	})
	// 设置flat后可以用这个修改fixed区域的表格线
	//m.proxyLogsGrid.SetFixedGridLineColor(colors.ClRed)
	//m.proxyLogsGrid.SetAnchors(types.NewSet(types.AkBottom, types.AkRight))
	//m.proxyLogsGrid.SetVisible(false)

	//绘制表格头
	m.proxyLogsGridHead()

	m.proxyLogsGrid.SetOnButtonClick(func(sender lcl.IObject, aCol, aRow int32) {
		logGridSelRowIndex = m.proxyLogsGrid.RowCount() - aRow
		m.selectGridUpdate()
	})
	m.proxyLogsGrid.SetOnDblClick(func(sender lcl.IObject) {
		m.selectGridUpdate()
	})

	//m.proxyLogsGrid.SetOnSetEditText(m.onGridSetEditText)
	//选中列表某行列数据
	m.proxyLogsGrid.SetOnSelectCell(func(sender lcl.IObject, aCol, aRow int32, canSelect *bool) {
		logGridSelRowIndex = m.proxyLogsGrid.RowCount() - aRow
		if aCol == 1 {
			//fmt.Println("row:", aRow, *canSelect, m.proxyLogsGrid.RowCount()-aRow)
			//放到剪切版
			logGridSelCol2Value = m.proxyLogsGrid.Cells(aCol, logGridSelRowIndex)
		}
	})
	//设置初始行数 1
	m.proxyLogsGrid.SetRowCount(m.proxyLogsGridRow)

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
		if logGridSelRowIndex != -1 {
			if row, ok := m.ProxyDetails[logGridSelRowIndex]; ok {
				if d, err := json.Marshal(row); err == nil {
					lcl.Clipboard.SetAsText(string(d))
				}
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
	if rowData, ok := m.ProxyDetails[logGridSelRowIndex]; ok {
		m.ProxyDetailUI.RequestDetailViewPanel.updateRequestSheet(rowData)
		m.ProxyDetailUI.RequestDetailViewPanel.updateResponseSheet(rowData)
		//if d, err := json.Marshal(rowData); err == nil {
		//	fmt.Println(" row", logGridSelRowIndex, "proxyDetail:", rowData.TargetUrl, " JSON:", string(d))
		//}
	}
}

//代理日志grid表格头
func (m *TGUIForm) proxyLogsGridHead() {
	var colNo = m.proxyLogsGrid.Columns().Add()
	colNo.SetWidth(100)
	colNo.Title().SetCaption("序号")
	colNo.Title().SetAlignment(types.TaCenter)
	colNo.SetAlignment(types.TaCenter)
	colNo.SetReadOnly(true)
	//col1.SetReadOnly(true)

	var colAddr = m.proxyLogsGrid.Columns().Add()
	colAddr.SetWidth(uiWidth - 180)
	colAddr.Title().SetCaption("地址 - (右键菜单复制)")
	colAddr.Title().SetAlignment(types.TaCenter)
	colAddr.SetReadOnly(true)

	var colDetailBtn = m.proxyLogsGrid.Columns().Add()
	colDetailBtn.SetWidth(60)
	colDetailBtn.SetButtonStyle(types.CbsButtonColumn)
	colDetailBtn.Title().SetCaption("详情")
	colDetailBtn.Title().SetAlignment(types.TaCenter)
	colDetailBtn.SetAlignment(types.TaCenter)
}

var (
	logGridSelCol2Value string      //选中表格第二列的值
	logGridSelRowIndex  int32  = -1 //选中表格行下标
	logGridInsertRow    int32  = 1  //在第指定行插入
)

//代理日志grid添加一行
func (m *TGUIForm) proxyLogsGridAdd(proxyDetail *entity.ProxyDetail) {
	lcl.ThreadSync(func() {
		//在指定行插入行
		m.proxyLogsGrid.InsertColRow(false, logGridInsertRow)
		//给指定行的列设置值
		m.proxyLogsGrid.SetCells(0, logGridInsertRow, fmt.Sprintf("%v", proxyDetail.ID))
		m.proxyLogsGrid.SetCells(1, logGridInsertRow, proxyDetail.Method+" - "+proxyDetail.TargetUrl)
		m.proxyLogsGrid.SetCells(2, logGridInsertRow, "查看")
		//计算增加一行
		atomic.AddInt32(&m.proxyLogsGridRow, 1)
		var r = atomic.LoadInt32(&m.proxyLogsGridRow)
		//给表格设置新总行数
		m.proxyLogsGrid.SetRowCount(r)
	})
}

//代理日志grid清除-还未完善
func (m *TGUIForm) proxyLogsGridClear() {
	m.proxyLogsGrid.Clear()
	logGridSelCol2Value = ""
	logGridSelRowIndex = -1
	m.proxyLogsGridRow = 1
	m.proxyLogsGrid.SetRowCount(m.proxyLogsGridRow)
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
	//更新集合内容
	m.ProxyDetails[proxyDetail.ID] = proxyDetail
	m.selectGridUpdate()
}
