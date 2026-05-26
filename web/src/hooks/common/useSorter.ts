import { ref } from 'vue';
import { DataTableSortState } from 'naive-ui';
import { SorterMultiple } from 'naive-ui/es/data-table/src/interface';

export default function useSorter(callback: Function = function () {}) {
  const sortStatesRef = ref<DataTableSortState[]>([]);

  function updateSorter(sorter: DataTableSortState | DataTableSortState[] | null) {
    if (sorter === null) {
      // 默认排序
      sortStatesRef.value = [];
    } else if (Array.isArray(sorter)) {
      // 多字段排序
      sortStatesRef.value = filterMultipleSorters(sorter);
    } else {
      // 单字段排序
      sortStatesRef.value = [];
      sortStatesRef.value.push(sorter);
    }
    callback();
  }

  // 多字段排序，将不参与排序的字段过滤掉，再根据设置的优先级对排序器重新排列
  function filterMultipleSorters(sorters: DataTableSortState[]): DataTableSortState[] {
    const filter = sorters.filter((item) => item.order !== false);
    return filter.sort((a, b) => {
      if (typeof a.sorter === 'object' && typeof b.sorter === 'object') {
        const aSorter = a.sorter as SorterMultiple;
        const bSorter = b.sorter as SorterMultiple;
        return bSorter.multiple - aSorter.multiple;
      }
      return 0;
    });
  }

  return {
    updateSorter,
    sortStatesRef,
  };
}
