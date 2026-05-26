import { renderIcon } from '@/utils';
import * as antdIcons from '@vicons/antd';
// import * as x5Icons from '@vicons/ionicons5';

export const constantRouterIcon = {};

export function createRouterIcon() {
  for (const element of Object.keys(antdIcons)) {
    constantRouterIcon[element] = renderIcon(antdIcons[element]);
  }

  // for (const element of Object.keys(x5Icons)) {
  //   constantRouterIcon[element] = renderIcon(x5Icons[element]);
  // }
}
