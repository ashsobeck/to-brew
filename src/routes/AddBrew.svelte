<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew } from '$lib/types';

	const addBrew = (e: Event) => {
		let newBrew: ToBrew = {
			id: crypto.randomUUID(),
			name: 'new brew',
			bean: '',
			done: false,
			time: new Date(),
			data: {}
		};

		tobrews.update((brews) => [newBrew, ...brews]);
	};
</script>

<button on:click={addBrew}>+ Add Brew</button>
