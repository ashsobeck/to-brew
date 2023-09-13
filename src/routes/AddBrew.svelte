<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew, Brew } from '$lib/types';

	let showBrew = false;
	//        const newBrew: ToBrew = {
	//            id: brew.id,
	//           name: brew.name,
	//         bean: brew.bean,
	//       brewed: false,
	//     time: new Date(brew.timeToBrew)
	//};

	//tobrews.update((brews) => [newBrew, ...brews]);
</script>

{#if showBrew}
	<form action="?/create">
		<div class="card p-2">
			<label class="label">Bean: <input class="input" name="bean" /> </label>
			<label class="label">Time: <input class="input" name="time" /> </label>
			<label class="label">Name: <input class="input" name="name" /> </label>
			<label class="label">Roaster: <input class="input" name="roaster" /> </label>
			<label class="label">Link: <input class="input" name="link" /> </label>
		</div>
	</form>
{/if}
<button disabled={showBrew} on:click={() => (showBrew = true)}>+ Add Brew</button>
