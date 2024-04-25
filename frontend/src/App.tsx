import { useEffect, useState } from "react";
import { createId } from "./utils/id";

interface QuoteType {
	id: string;
	author: string;
	quote: string;
}

function App() {
	const [quotes, setQuotes] = useState<QuoteType[] | null>([]);
	const [author, setAuthor] = useState<string>("");
	const [quote, setQuote] = useState<string>("");

	async function getQuotes() {
		const response = await fetch("http://localhost:8080/quotes");
		const data = await response.json();
		if (!data) {
			return;
		}

		setQuotes(data);
	}

	async function handleSubmit() {
		if (!quote || !author) {
			alert("Author and quote must be valid values.");
			return;
		}

		const id = createId();

		const response = await fetch("http://localhost:8080/quotes", {
			method: "POST",
			body: JSON.stringify({ id: id, author: author, quote: quote }),
		});
		const data = await response.json();

		console.log(data);

		setQuote("");
		setAuthor("");

		getQuotes();
	}

	useEffect(() => {
		getQuotes();
	}, []);

	return (
		<>
			<div className="max-w-2xl mx-auto py-16">
				<h1 className="text-4xl font-bold text-gray-800 text-center mb-8">
					Quotes
				</h1>
				<div className="flex flex-col">
					<p className="p-2 text-gray-800 font-medium">Add a quote</p>
					<div className="flex space-x-4">
						<input
							className="w-1/2 p-2"
							placeholder="Enter author..."
							value={author}
							onChange={(e) => setAuthor(e.target.value)}
						/>
						<input
							className="w-1/2 p-2 "
							placeholder="Enter quote..."
							value={quote}
							onChange={(e) => setQuote(e.target.value)}
						/>
					</div>
					<button
						className="bg-gray-800 text-white rounded-md p-2 hover:bg-gray-600 mt-8"
						onClick={handleSubmit}
					>
						Submit
					</button>
				</div>
				<div className="flex flex-col space-y-4 justify-center mt-8">
					{quotes &&
						quotes.map((quotes) => (
							<div key={quotes.id} className="p-2">
								<div className="flex">
									<p className="mb-2">{quotes.quote}</p>
								</div>
								<div>
									<p> - {quotes.author}</p>
								</div>
							</div>
						))}
				</div>

				{!quotes ||
					(quotes.length == 0 && (
						<p className="text-sm text-gray-800 font-medium text-center">
							Server not running or there is no quotes available
							to show...
						</p>
					))}
			</div>
		</>
	);
}

export default App;
