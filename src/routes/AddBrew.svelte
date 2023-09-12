<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew, Brew } from '$lib/types';

	let showBrew = false;
	const addBrew = async () => {
		const response = await fetch('http://localhost:3333/tobrews/', {
			method: 'POST',
			body: JSON.stringify({
				name: 'new brew',
				roaster: { String: 'S&W', Valid: true },
				link: { String: 'https://us.mystery.coffee', Valid: true },
				brewed: false,
				timeToBrew: new Date().toISOString()
			}),
			headers: {
				'Content-Type': 'application/json',
				Origin: 'http://localhost:5173/'
			}
		});

		const brew = (await response.json()) as Brew;
		console.log(brew);

		let newBrew: ToBrew = {
			id: brew.id,
			name: 'new brew',
			bean: '',
			brewed: false,
			time: new Date()
		};

		tobrews.update((brews) => [newBrew, ...brews]);
	};
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
<button on:click={() => (showBrew = true)}>+ Add Brew</button>
