<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew, Brew } from '$lib/types';

	const addBrew = async () => {
		const response = await fetch('http://localhost:3333/tobrews/', {
			method: 'POST',
			body: JSON.stringify({
				name: 'new brew',
				roaster: 'S&W',
				link: 'https://us.mystery.coffee',
				date: new Date()
			}),
			headers: {
				'Content-Type': 'application/json',
				Origin: 'http://localhost:5173/'
			}
		});

		const brew = (await response.json()) as Brew;

		let newBrew: ToBrew = {
			id: brew.id,
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
