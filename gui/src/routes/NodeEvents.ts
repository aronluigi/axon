import { availableNodes, nodesStore } from "$lib";
import { type Node } from '@xyflow/svelte';
import { get } from "svelte/store";
import { UpdateFlowState, type InputMaybe, type InputNode, type InputPort, type UpdateFlowStateNodesParam } from "../generated/graphql";

const dataTransferFormat: string = 'axon/flow'

export const onNodeDragStart = (event: DragEvent, uuid: string) => {
  if (!event.dataTransfer) {
    return
  }

  event.dataTransfer.setData(dataTransferFormat, uuid)
  event.dataTransfer.effectAllowed = 'move'
}

export const onNodeDragOver = (event: DragEvent) => {
  event.preventDefault();

  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = 'move';
  }
};

export const onNodeDrop = (screenToFlowPosition: any) => (event: DragEvent) => {
  event.preventDefault()

  if (!event.dataTransfer) {
    return
  }

  const uuid = event.dataTransfer.getData(dataTransferFormat)
  const an = get(availableNodes)

  if (!an[uuid]) {
    return
  }

  const position = screenToFlowPosition({
    x: event.clientX,
    y: event.clientY
  })

  const newNode: Node = {
    id: new Date().valueOf().toString(),
    type: 'custom',
    position: position,
    data: an[uuid]
  }

  const state = get(nodesStore)
  state.push(newNode)

  const pl = state.map(x => ({
    id: x.id,
    position: x.position,
    type: x.type,
    data: {
      uuid: x.data.uuid,
      displayName: x.data.label,
      packageName: x.data.package,
      inputPorts: x.data.inPorts as InputPort[],
      outputPorts: x.data.outPorts as InputPort[],
    } as InputNode
  } as UpdateFlowStateNodesParam))
  console.log(pl)
  UpdateFlowState({ variables: { nodes: pl } })
    .then(r => {
      if (r.errors) {
        console.log(r.errors)
        return
      }

      nodesStore.set(state)
    })
    .catch(e => {
      throw new Error(e)
    })
}
