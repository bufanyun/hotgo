/**
   * 控制表单列数布局
   * @param layout 列布局: 1,2,3,4
   */
  const formColLayout = function (layout, colSetting, fixedSetting) {
    if (layout && layout === '1') {
      // 1列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 24, sm: 24, lg: 24, xl: 24 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 6, sm: 6, lg: 6, xl: 6 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 18, sm: 18, lg: 18, xl: 18 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 6, sm: 6, lg: 6, xl: 6 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 18, sm: 18, lg: 18, xl: 18 }
        }
      }
    } else if (layout && layout === '2') {
      // 2列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 12, sm: 12, lg: 12, xl: 12 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 8, sm: 8, lg: 8, xl: 8 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 16, sm: 16, lg: 16, xl: 16 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 4, sm: 4, lg: 4, xl: 4 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 20, sm: 20, lg: 20, xl: 20 }
        }
      }
    } else if (layout && layout === '3') {
      // 3列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 8, sm: 8, lg: 8, xl: 8 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 9, sm: 9, lg: 9, xl: 9 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 15, sm: 15, lg: 15, xl: 15 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 3, sm: 3, lg: 3, xl: 3 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 21, sm: 21, lg: 21, xl: 21 }
        }
      }
    } else if (layout && layout === '4') {
      // 4列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 6, sm: 6, lg: 6, xl: 6 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 8, sm: 8, lg: 8, xl: 8 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 16, sm: 16, lg: 16, xl: 16 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 2, sm: 2, lg: 2, xl: 2 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 22, sm: 22, lg: 22, xl: 22 }
        }
      }
    } else {
      // 默认1列
      return {
        cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
        labelCol: { xs: 4, sm: 4, lg: 4, xl: 4 },
        wrapperCol: { xs: 20, sm: 20, lg: 20, xl: 20 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol: { xs: 4, sm: 4, lg: 4, xl: 4 },
          wrapperCol: { xs: 20, sm: 20, lg: 20, xl: 20 }
        }
      }
    }
  }

  /**
   * 控制详情页表单列数布局
   * @param layout 列布局: 1,2,3,4
   */
  const flowformColLayout = function (layout, colSetting, fixedSetting) {
    if (layout && layout === '1') {
      // 1列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 24, sm: 24, lg: 24, xl: 24 },
        labelCol:
          colSetting && colSetting.labelCol
            ? colSetting.labelCol
            : { xs: 22, sm: 22, lg: 22, xl: 22 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 22, sm: 22, lg: 22, xl: 22 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 24, sm: 24, lg: 24, xl: 24 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 20, sm: 20, lg: 20, xl: 20 }
        }
      }
    } else if (layout && layout === '2') {
      // 2列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 12, sm: 12, lg: 12, xl: 12 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 6, sm: 6, lg: 6, xl: 6 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 18, sm: 18, lg: 18, xl: 18 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 3, sm: 3, lg: 3, xl: 3 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 21, sm: 21, lg: 21, xl: 21 }
        }
      }
    } else if (layout && layout === '3') {
      // 3列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 8, sm: 8, lg: 8, xl: 8 },
        labelCol:
          colSetting && colSetting.labelCol ? colSetting.labelCol : { xs: 6, sm: 6, lg: 6, xl: 6 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 18, sm: 18, lg: 18, xl: 18 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 2, sm: 2, lg: 2, xl: 2 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 22, sm: 22, lg: 22, xl: 22 }
        }
      }
    } else if (layout && layout === '4') {
      // 4列
      return {
        cols: colSetting && colSetting.cols ? colSetting.cols : { xs: 6, sm: 6, lg: 6, xl: 6 },
        labelCol:
          colSetting && colSetting.labelCol
            ? colSetting.labelCol
            : { xs: 24, sm: 24, lg: 24, xl: 24 },
        wrapperCol:
          colSetting && colSetting.wrapperCol
            ? colSetting.wrapperCol
            : { xs: 24, sm: 24, lg: 24, xl: 24 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol:
            fixedSetting && fixedSetting.labelCol
              ? fixedSetting.labelCol
              : { xs: 24, sm: 24, lg: 24, xl: 24 },
          wrapperCol:
            fixedSetting && fixedSetting.wrapperCol
              ? fixedSetting.wrapperCol
              : { xs: 24, sm: 24, lg: 24, xl: 24 }
        }
      }
    } else {
      // 默认1列
      return {
        cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
        labelCol: { xs: 4, sm: 4, lg: 4, xl: 4 },
        wrapperCol: { xs: 20, sm: 20, lg: 20, xl: 20 },
        fixed: {
          cols: { xs: 24, sm: 24, lg: 24, xl: 24 },
          labelCol: { xs: 4, sm: 4, lg: 4, xl: 4 },
          wrapperCol: { xs: 20, sm: 20, lg: 20, xl: 20 }
        }
      }
    }
  }
  /**
   * 控制弹窗modal的宽高
   * @param widthRatio 弹窗modal宽度占屏幕比例
   * @param heightRatio 弹窗modal高度占屏幕比例
   * @param cutHeight 弹窗modal的body的高度比整个弹窗modal少的高度
   * @param bodyOverflowY 弹窗modal的body的overflow-y属性
   */
  const modalWidthAndHeight = function (
    widthRatio = 0.7,
    heightRatio = 0.7,
    cutHeight = 40,
    bodyOverflowY = 'auto',
    backgroundColor = '#fff'
  ) {
    const modalStyleWidth = document.documentElement.clientWidth * widthRatio + 'px'
    // let modalStyleHeight = document.documentElement.clientHeight * heightRatio + 'px';
    const bodyStyleWidth = document.documentElement.clientWidth * widthRatio + 'px'
    const bodyStyleHeight = document.documentElement.clientHeight * heightRatio - cutHeight + 'px'
    return {
      modalStyle: {
        // 弹窗宽高控制
        width: modalStyleWidth,
        height: 'auto'
      },
      bodyStyle: {
        // 弹窗body宽高控制
        width: bodyStyleWidth,
        height: bodyStyleHeight,
        overflowY: bodyOverflowY,
        backgroundColor
      }
    }
  }
  /**
   * 根据布局列数控制弹窗modal的宽高
   * （单列： 580*420  适用于1~6项内容时、
   *   两列：960*500   适用于6~12项内容时、
   *   三列：1280*600） 适用于12项以上
   * @param layout 列数
   * @param cutHeight 弹窗modal的body的高度比整个弹窗modal少的高度
   * @param bodyOverflowY 弹窗modal的body的overflow-y属性
   */
  const modalWidthAndHeightBylayout = function (
    layout = 1,
    cutHeight = 40,
    bodyOverflowY = 'auto',
    backgroundColor = '#fff'
  ) {
    let modalStyleWidth = null
    // let modalStyleHeight = document.documentElement.clientHeight * heightRatio + 'px';
    let bodyStyleWidth = null
    let bodyStyleHeight = null
    let bodyStylePadding = null
    if (layout === 1) {
      modalStyleWidth = 580 + 'px'
      bodyStyleWidth = 580 + 'px'
      bodyStyleHeight = 420 - cutHeight + 'px'
      bodyStylePadding = '20px 50px'
    } else if (layout === 2) {
      modalStyleWidth = 960 + 'px'
      bodyStyleWidth = 960 + 'px'
      bodyStyleHeight = 500 - cutHeight + 'px'
      bodyStylePadding = '20px'
    } else if (layout === 3) {
      modalStyleWidth = 1280 + 'px'
      bodyStyleWidth = 1280 + 'px'
      bodyStyleHeight = 600 - cutHeight + 'px'
      bodyStylePadding = '20px 40px'
    }
    return {
      modalStyle: {
        // 弹窗宽高控制
        width: modalStyleWidth,
        height: 'auto'
      },
      bodyStyle: {
        // 弹窗body宽高控制
        width: bodyStyleWidth,
        height: bodyStyleHeight,
        padding: bodyStylePadding,
        overflowY: bodyOverflowY,
        backgroundColor
      }
    }
  }
  /**
   * 获取内容高度: 传入元素的父元素高度 - 父元素除去内容之外的元素的高度
   * @param wrapId 元素的id
   * @param cutHeight 内容与父元素相差的高度,一般包括:父元素padding,按钮高度及margin,及其他
   */
  const getContentHeight = function (params) {
    // 默认值auto
    let contentHeight = 'auto'
    if (params) {
      // 获取元素父元素
      const el = document.querySelector('#' + params.wrapId)
      if (el) {
        // 获取父元素高度
        const parentHeight = el.parentNode.offsetHeight
        // 传入的元素高度设为父元素高度
        // el.style.height = parentHeight + 'px'; 为加resize改变大小,注释掉这句.
        // 内容高度为 父元素高度 减去 父元素除去内容之外的元素的高度
        if (params.cutHeight) {
          contentHeight = parentHeight - params.cutHeight + 'px'
        } else {
          contentHeight = parentHeight + 'px'
        }
      }
    }
    return contentHeight
  }

  /**
   * 校验不通过,滚动到校验不通过元素位置,提示用户
   */
  const getFirstCheckErrorElement = function (err) {
    for (const key in err) {
      // return document.querySelector("div[for='" + key + "']") || document.querySelector("div[for='" + key + "Alias']") || document.querySelector("div[for='" + key + "Name']");
      return document.querySelector('#' + key)
    }
  }
 const appointModalWidthAndHeight = function (
    height,
    width,
    cutHeight,
    bodyOverflowY = 'auto',
    backgroundColor = '#fff'
  ) {
     let modalStyleWidth
     let bodyStyleWidth
     let bodyStyleHeight
    if (width === '100%' && height === '100%') {
      // 全屏模式
       modalStyleWidth = document.documentElement.clientWidth + 'px'
       bodyStyleWidth = document.documentElement.clientWidth + 'px'
       bodyStyleHeight = document.documentElement.clientHeight - cutHeight + 'px'
    } else {
      // 指定宽高
       modalStyleWidth = width + 'px'
       bodyStyleWidth = width + 'px'
       bodyStyleHeight = height - cutHeight + 'px'
    }
    return {
      modalStyle: {
        // 弹窗宽高控制
        width: modalStyleWidth,
        height: 'auto'
      },
      bodyStyle: {
        // 弹窗body宽高控制
        width: bodyStyleWidth,
        height: bodyStyleHeight,
        overflowY: bodyOverflowY,
        backgroundColor
      }
    }
  }
  export {
    formColLayout,
    flowformColLayout,
    modalWidthAndHeight,
    modalWidthAndHeightBylayout,
    getContentHeight,
    getFirstCheckErrorElement,
    appointModalWidthAndHeight
  }
