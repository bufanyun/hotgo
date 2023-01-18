import { renderIcon } from '@/utils/index';
import * as antdIcons from '@vicons/antd';
// import * as x5Icons from '@vicons/ionicons5';

export const constantRouterIcon = {};
for (const element of Object.keys(antdIcons)) {
  constantRouterIcon[element] = renderIcon(antdIcons[element]);
}

// for (const element of Object.keys(x5Icons)) {
//   constantRouterIcon[element] = renderIcon(x5Icons[element]);
// }
