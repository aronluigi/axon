import { writable, type Writable } from 'svelte/store';
import { type Node } from '@xyflow/svelte';
import { FetchAvailableNodes } from '../../generated/graphql';

const testNodes: Node[] = [
  {
    id: '1', // required and needs to be a string
    position: { x: 0, y: 0 }, // required
    data: { label: 'hey' } // required
  },
  {
    id: '2',
    position: { x: 100, y: 100 },
    data: { label: 'world' }
  }
];


export const nodesStore: Writable<Node[]> = writable([])

FetchAvailableNodes({}).subscribe((s) => {
  if (Object.keys(s.data).length === 0) {
    nodesStore.set([])
    return
  }

  let res: Node[] = s.data.nodes.map((x) => {
    return {
      id: x.displayName,
      type: 'custom',
      position: { x: 100, y: 100 },
      data: {
        label: x.displayName,
        package: x.packageName,
        inPorts: x.inputPorts,
        outPorts: x.outputPorts
      }
    } as Node
  })

  nodesStore.set(testNodes.concat(res))
})
