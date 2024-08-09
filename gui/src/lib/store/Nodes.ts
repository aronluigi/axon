import { writable, type Writable } from 'svelte/store';
import { type Node } from '@xyflow/svelte';
import { FetchAvailableNodes } from '../../generated/graphql';
import { type NodeObj } from '$lib/types';

export const nodesStore: Writable<Node[]> = writable([])
export const availableNodes: Writable<{ [uuid: string]: NodeObj }> = writable({})

FetchAvailableNodes({}).subscribe((s) => {
  if (Object.keys(s.data).length === 0) {
    nodesStore.set([])
    return
  }

  const nodes: { [uuid: string]: NodeObj } = {}
  s.data.nodes.forEach(x => {
    nodes[x.uuid] = {
      uuid: x.uuid,
      label: x.displayName,
      package: x.packageName,
      inPorts: x.inputPorts,
      outPorts: x.outputPorts
    } as NodeObj
  });

  availableNodes.set(nodes)
})
