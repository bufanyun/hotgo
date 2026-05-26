import { ref, Ref } from "vue";

type Resettable =
  /**
   * 基础形式为一个元组，包含：
   * 1. `Ref<boolean>`：状态引用，可以直接访问和修改状态值
   * 2. `() => void`：切换真假
   * 3. `(value: boolean) => void`：设置状态
   * 4.`() => void`：设置为`true`方法
   * 5.`() => void`：设置为`false`方法
   */
  [Ref<boolean>, () => void, () => void] & /**
   * 同时扩展为一个对象形式，提供以下属性和方法：
   * - `loading: Ref<T>`：状态引用，与元组中的第一个元素相同
   * - `startLoading: () => void`：开始加载方法，与元组中的第二个元素相同
   * - `endLoading: () => void`：结束加载方法，与元组中的第三个元素相同
   */ {
    /** 加载状态引用 */
    bool: Ref<boolean>;
    toggle: () => void;
    setBool: (value: boolean) => void;
    setTrue: () => void;
    setFalse: () => void;
  };
export default function useBoolean(initValue = false): Resettable {
  const bool = ref(initValue);

  function setBool(value: boolean) {
    bool.value = value;
  }
  function setTrue() {
    setBool(true);
  }
  function setFalse() {
    setBool(false);
  }
  function toggle() {
    setBool(!bool.value);
  }

  // 返回带加载状态引用、开始加载和结束加载方法的扩展数组
  return (Object.assign([bool, toggle, setBool, setTrue, setFalse], {
    bool,
    toggle,
    setBool,
    setTrue,
    setFalse,
  }) as unknown) as Resettable;
}
