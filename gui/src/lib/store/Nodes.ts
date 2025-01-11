import { writable } from 'svelte/store';
import { type Node } from '@xyflow/svelte';
import { FetchAvailableNodes } from '../../generated/graphql';
import { type NodeObj, type NodeList } from '$lib/types';

export const nodesStore = writable<Node[]>([])
export const availableNodes = writable<NodeList>({})

export const getAvailableNodes = async () => {
  let getNodes: () => Promise<NodeList> = () => new Promise((resolve) => {
    const nodes: NodeList = {}

    FetchAvailableNodes({}).subscribe((s) => {
      if (Object.keys(s.data).length === 0) {
        return []
      }

      s.data.nodes.forEach(x => {
        nodes[x.uuid] = {
          uuid: x.uuid,
          label: x.displayName,
          package: x.packageName,
          inPorts: x.inputPorts,
          outPorts: x.outputPorts
        } as NodeObj
      });

      resolve(nodes)
    })
  })

  return await getNodes()
}
