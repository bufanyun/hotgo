export default {
  name: 'VueFragment',
  functional: true,
  render (h, ctx) {
    return ctx.children.length > 1 ? h('div', {}, ctx.children) : ctx.children
  }
}
