class RibbonData {
  background: string
  outline: string
  constructor(
    background: string,
    outline : string
  ) {
    this.background = background;
    this.outline = outline;
  }
}

const ribbonData = {
  julie: new RibbonData('#DCB0F2', '#a35ec9'),
  champion: new RibbonData('#66C5CC', '#358d95'),
  battle: new RibbonData('#F89C74', '#754834'),
  contest: new RibbonData('#FE88B1', '#8e3d5a'),
  stats: new RibbonData('#9EB9F3', '#62749b'),
  shopping: new RibbonData('#F6CF71', '#d5a32c')
} as any

export default ribbonData
export type { RibbonData }