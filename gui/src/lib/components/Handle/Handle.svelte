<script lang="ts">
	import { Handle, Position } from '@xyflow/svelte';
	import { PortType } from '$lib/types';

	interface Props {
		portType: PortType;
		id: string;
	}

	let { portType, id }: Props = $props();

	const position = portType === PortType.in ? Position.Left : Position.Right;
	const handleType = portType === PortType.in ? 'target' : 'source';
	const portColor = portType === PortType.in ? 'bg-red-700' : 'bg-green-700';
	const style = `
    width: 1.25rem;
    height: 1.25rem;
    position: absolute;
    top: 50%;
    ${position === Position.Left ? 'left' : 'right'}: 50%;
    opacity: 0;
  `;
</script>

<div class={`relative inline-block w-5 h-5 rounded-full ${portColor}`}>
	<Handle type={handleType} {position} {style} {id} isConnectable={true} />
</div>
