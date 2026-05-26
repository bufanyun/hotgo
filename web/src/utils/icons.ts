import { Component } from 'vue';
import { renderIcon } from '@/utils/index';
import * as antdIcons from '@vicons/antd';
// import * as x5Icons from '@vicons/ionicons5';

export const Icons = {};
for (const element of Object.keys(antdIcons)) {
  Icons[element] = renderIcon(antdIcons[element]);
}

// for (const element of Object.keys(x5Icons)) {
//   constantRouterIcon[element] = renderIcon(x5Icons[element]);
// }

export function getIconComponent(name: string): Component {
  return Icons[name];
}
