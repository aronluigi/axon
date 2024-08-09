import { availableNodes, nodesStore } from "$lib";
import { useSvelteFlow, type XYPosition } from "@xyflow/svelte"
import { type Node } from '@xyflow/svelte';
import { getContext } from "svelte";
import { get } from "svelte/store";

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

  nodesStore.update((n) => [...n, newNode])
}
