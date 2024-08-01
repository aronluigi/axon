export type NodePort = {
  name: string;
  type: string;
  pkgPath: string;
};

export type Node = {
  data: {
    label: string;
    package: string;
    inPorts: NodePort[];
    outPorts: NodePort[];
  };
};
