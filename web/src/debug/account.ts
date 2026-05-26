import { isArray } from "@/utils/is";

interface Account {
  /**
   * 账号名称
   */
  name: string;
  /**
   * 用户名
   */
  username: string;
  /**
   * 密码
   */
  password: string;
}
/**
 * 获取配置的账号信息
 * @returns {[]Account} 返回账号信息数组
 */
export function getDemoAccounts() {
  
  let envConf = import.meta.env.VITE_APP_DEMO_ACCOUNT || "";
  // 帐号密码一样
  // [["username"],["username","password"],["username","password","name"]]
  try {
    let accounts = JSON.parse(envConf);
    if (accounts && isArray(accounts)) {
      return accounts.map((item: String[]) => {
        let [username = "", password = "", name = ""] = item;
        username = username;
        password = password || username;
        name = name || username;
        return {
          name,
          username,
          password,
        } as Account;
      });
    }
  } catch (error) {}
  return [] as Account[];
}
