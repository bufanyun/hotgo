// 关键词配置
interface IKeywordOption {
  keyword: string | RegExp;
  color?: string;
  bgColor?: string;
  style?: Record<string, any>;
  // 高亮标签名
  tagName?: string;
  // 忽略大小写
  caseSensitive?: boolean;
  // 自定义渲染高亮html
  // eslint-disable-next-line no-unused-vars
  renderHighlightKeyword?: (content: string) => any;
}

type IKeyword = string | IKeywordOption;

export interface IMatchIndex {
  index: number;
  subString: string;
}

// 关键词索引
export interface IKeywordParseIndex {
  keyword: string | RegExp;
  indexList: IMatchIndex[];
  option?: IKeywordOption;
}

// 关键词
export interface IKeywordParseResult {
  start: number;
  end: number;
  subString?: string;
  option?: IKeywordOption;
}

// 计算
const getKeywordIndexList = (content: string, keyword: string | RegExp, flags = 'ig') => {
  const reg = new RegExp(keyword, flags);
  const res = (content as any).matchAll(reg);
  const arr = [...res];
  const allIndexArr: IMatchIndex[] = arr.map((e) => ({
    index: e.index,
    subString: e['0'],
  }));
  return allIndexArr;
};

// 驼峰转换横线
function humpToLine(name: string) {
  return name.replace(/([A-Z])/g, '-$1').toLowerCase();
}

const renderNodeTag = (subStr: string, option: IKeywordOption) => {
  const s = subStr;
  if (!option) {
    return s;
  }
  const { tagName = 'mark', bgColor, color, style = {}, renderHighlightKeyword } = option;
  if (typeof renderHighlightKeyword === 'function') {
    return renderHighlightKeyword(subStr);
  }
  style.backgroundColor = bgColor;
  style.color = color;

  const styleContent = Object.keys(style)
    .map((k) => `${humpToLine(k)}:${style[k]}`)
    .join(';');
  const styleStr = `style="${styleContent}"`;
  return `<${tagName} ${styleStr}>${s}</${tagName}>`;
};

const renderHighlightHtml = (content: string, list: any[]) => {
  let str = '';
  list.forEach((item) => {
    const { start, end, option } = item;
    const s = content.slice(start, end);
    const subStr = renderNodeTag(s, option);
    str += subStr;
    item.subString = subStr;
  });
  return str;
};

// 解析关键词为索引
const parseHighlightIndex = (content: string, keywords: IKeyword[]) => {
  const result: IKeywordParseIndex[] = [];
  keywords.forEach((keywordOption: IKeyword) => {
    let option: IKeywordOption = { keyword: '' };
    if (typeof keywordOption === 'string') {
      option = { keyword: keywordOption };
    } else {
      option = keywordOption;
    }
    const { keyword, caseSensitive = true } = option;
    const indexList = getKeywordIndexList(content, keyword, caseSensitive ? 'g' : 'gi');
    const res = {
      keyword,
      indexList,
      option,
    };
    result.push(res);
  });
  return result;
};

const parseHighlightString = (content: string, keywords: IKeyword[]) => {
  const result = parseHighlightIndex(content, keywords);
  const splitList: IKeywordParseResult[] = [];
  const findSplitIndex = (index: number, len: number) => {
    for (let i = 0; i < splitList.length; i++) {
      const cur = splitList[i];
      // 有交集
      if (
        (index > cur.start && index < cur.end) ||
        (index + len > cur.start && index + len < cur.end) ||
        (cur.start > index && cur.start < index + len) ||
        (cur.end > index && cur.end < index + len) ||
        (index === cur.start && index + len === cur.end)
      ) {
        return -1;
      }
      // 没有交集，且在当前的前面
      if (index + len <= cur.start) {
        return i;
      }
      // 没有交集，且在当前的后面的，放在下个迭代处理
    }
    return splitList.length;
  };
  result.forEach(({ indexList, option }: IKeywordParseIndex) => {
    indexList.forEach((e) => {
      const { index, subString } = e;
      const item = {
        start: index,
        end: index + subString.length,
        option,
      };
      const splitIndex = findSplitIndex(index, subString.length);
      if (splitIndex !== -1) {
        splitList.splice(splitIndex, 0, item);
      }
    });
  });

  // 补上没有匹配关键词的部分
  const list: IKeywordParseResult[] = [];
  splitList.forEach((cur, i) => {
    const { start, end } = cur;
    const next = splitList[i + 1];
    // 第一个前面补一个
    if (i === 0 && start > 0) {
      list.push({ start: 0, end: start, subString: content.slice(0, start) });
    }
    list.push({ ...cur, subString: content.slice(start, end) });
    // 当前和下一个中间补一个
    if (next?.start > end) {
      list.push({
        start: end,
        end: next.start,
        subString: content.slice(end, next.start),
      });
    }
    // 最后一个后面补一个
    if (i === splitList.length - 1 && end < content.length - 1) {
      list.push({
        start: end,
        end: content.length - 1,
        subString: content.slice(end, content.length),
      });
    }
  });
  return list;
};

// 生成关键词高亮的html字符串
const highHtml = (content: string, keywords: IKeyword[]) => {
  const splitList = parseHighlightString(content, keywords);
  return {
    highText: renderHighlightHtml(content, splitList),
    highList: splitList,
  };
};

export default highHtml;
