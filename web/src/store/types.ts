import { IAsyncRouteState } from '@/store/modules/asyncRoute';
import { IUserState } from '@/store/modules/user';
import { ILockscreenState } from '@/store/modules/lockscreen';
import { ITabsViewState } from '@/store/modules/tabsView';
import { INotificationStore } from '@/store/modules/notification';

export interface IStore {
  asyncRoute: IAsyncRouteState;
  user: IUserState;
  lockscreen: ILockscreenState;
  tabsView: ITabsViewState;
  notification: INotificationStore;
  count: number;
}
