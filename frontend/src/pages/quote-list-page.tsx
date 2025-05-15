import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import axios from "axios";
import { useState } from "react";

type Quote = {
  quote: string;
  author: string;
  book: string;
};

export const QuoteListPage = () => {
  const [quotes, setQuotes] = useState<Quote[]>([]);
  const [author, setAuthor] = useState("");
  const [book, setBook] = useState("");
  const [q, setQ] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  // Modal
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [randomQuote, setRandomQuote] = useState<Quote | null>(null);

  const fetchRandomQuote = async () => {
    const response = await axios.get("http://localhost:3000/quotes/random", {
      params: { author, book, q, page: currentPage },
    });

    setRandomQuote(response.data);
    setIsModalOpen(true);
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold mb-4">Quotes Library</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
        <Input
          placeholder="Filter by Author"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
        />
        <Input
          placeholder="Filter by Book"
          value={book}
          onChange={(e) => setBook(e.target.value)}
        />
        <Input
          placeholder="Filter by Quote"
          value={q}
          onChange={(e) => setQ(e.target.value)}
        />
      </div>

      <div className="flex justify-end mb-6">
        <Button onClick={fetchRandomQuote}>Get Random Quote</Button>
      </div>
    </div>
  );
};
