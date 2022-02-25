<script>
// 新加图标示例
const MtimeSvg = {
  template: `
            <svg t="1598838771152" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2725" xmlns:xlink="http://www.w3.org/1999/xlink" width="200" height="200"><path d="M511.573333 85.333333C276.053333 85.333333 85.333333 276.48 85.333333 512s190.72 426.666667 426.24 426.666667C747.52 938.666667 938.666667 747.52 938.666667 512S747.52 85.333333 511.573333 85.333333zM693.333333 689.066667 469.333333 554.666667 469.333333 298.666667l64 0 0 224 192 113.92L693.333333 689.066667z" p-id="2726"></path></svg>`
}
export default {
  name: 'MtimeIcon',
  template: `
            <a-icon :component="MtimeSvg" />
        `,
  data () {
    return {
      MtimeSvg
    }
  }
}
</script>
