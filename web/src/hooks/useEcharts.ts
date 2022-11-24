import * as echarts from 'echarts';
export default function useEcharts(dom: HTMLElement, theme?: string) {
  let instance = echarts.getInstanceByDom(dom);
  if (!instance) {
    instance = echarts.init(dom, theme);
  }
  return instance;
}
