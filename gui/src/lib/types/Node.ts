export type NodePort = {
  name: string
  type: string
  pkgPath: string
}

export type NodeObj = {
  uuid: string
  label: string
  package: string
  inPorts: NodePort[]
  outPorts: NodePort[]
}
