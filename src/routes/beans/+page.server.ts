
export async function load() {
	const res = await fetch(`http://127.0.0.1:3333/beans/`, {
		method: 'GET',
		headers: {
			Origin: 'http://127.0.0.1:5173'
		}
	});

	const beans = await res.json();
	return {
		beans: beans ?? []
	};
}

export const actions = {
	create: async ({ request }) => {
		const data = await request.formData();

		const res = await fetch(`http://127.0.0.1:3333/beans/`, {
			method: 'POST',
			body: JSON.stringify({
				name: data.get('name'),
				roaster: data.get('roaster'),
				country: data.get('country'),
				varietal: data.get('varietal'),
				process: data.get('process'),
				altitude: data.get('altitude'),
				notes: data.get('notes'),
				weight: parseFloat(data.get('weight'))
			}),
			headers: {
				Origin: 'http://127.0.0.1:5173'
			}
		});

		const newBean = await res.json();
		console.log(newBean);
		return { success: newBean !== null, bean: newBean };
	}
}
