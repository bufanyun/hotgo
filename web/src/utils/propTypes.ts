import { CSSProperties, VNodeChild } from 'vue';
import {
  createTypes,
  VueTypeExtendCallback,
  VueTypeValidableDef,
  VueTypesInterface,
} from 'vue-types';

export type VueNode = VNodeChild | JSX.Element;

type ExtendedPropTypes = VueTypesInterface & {
  readonly style: VueTypeValidableDef<CSSProperties>;
  readonly VNodeChild: VueTypeValidableDef<VueNode>;
};

class CustomVueTypes
  extends (createTypes({
    func: undefined,
    bool: undefined,
    string: undefined,
    number: undefined,
    object: undefined,
    integer: undefined,
  }) as VueTypesInterface)
  implements ExtendedPropTypes
{
  static extend(types: VueTypeExtendCallback[]): CustomVueTypes {
    return types.reduce((result, { name, ...callbacks }) => {
      result[name] = { getter: true, ...callbacks };
      return result;
    }, this);
  }

  readonly style!: VueTypeValidableDef<CSSProperties>;
  readonly VNodeChild!: VueTypeValidableDef<VueNode>;
}

const propTypes = CustomVueTypes.extend([
  {
    name: 'style',
    type: [String, Object],
    default: undefined,
  },
  {
    name: 'VNodeChild',
    type: undefined,
  },
]);

export { propTypes };
