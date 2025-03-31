import React, { useState, useEffect } from 'react';
import Header from '@/components/Header';
import Hero from '@/components/Hero';
import AnimatedBackground from '@/components/AnimatedBackground';
import TagFilter from '@/components/TagFilter';
import EnhancedPromptCard from '@/components/EnhancedPromptCard';
import { icons } from 'lucide-react';
import supabase from '../utils/supabase';
import { useToast } from '@/hooks/use-toast';
import { Button } from "@/components/ui/button";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"

export interface Prompt {
  id: string;
  title: string;
  description?: string;
  content?: string;
  like_count?: number;
  image_url?: string;
  created_at?: string;  // Using ISO 8601 format for timestamps
  updated_at?: string;
  deleted_at?: string;
  tags: string[];  // The array of tags
  trending: boolean;
}

const Index = () => {
  const { toast } = useToast();
  const [selectedTags, setSelectedTags] = useState<string[]>([]);
  const [selectedView, setSelectedView] = useState<string>("all");
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [filteredPrompts, setFilteredPrompts] = useState<Prompt[]>([]);
  const [page, setPage] = useState<number>(1);
  const [pageCount, setPageCount] = useState<number>(1);
  const [page_limit, setPageLimit] = useState<string>("10");
  const [totalCount, setTotalCount] = useState<number>(1);

  // Filter prompts whenever search query, tags, or view changes
  useEffect(() => {
    async function getPrompts() {
      try {
        const { data, error } = await supabase.rpc('get_prompts', {
          page: page,
          page_limit: Number(page_limit),
          search_tags: selectedTags.length > 0 ? selectedTags : undefined,
          is_trending: selectedView === "trending" ? true : false,
          search_query: searchQuery || undefined
        });
        console.log(data)
        if (error) {
          toast({
            title: 'Error',
            description: error.message
          })
        } else {
          setFilteredPrompts(data?.prompts ?? []);
          setPageCount(data?.total_pages ?? 1);
          setTotalCount(data?.total_count ?? 1);
        }
      } catch (e) {
        toast({
          title: 'Error',
          description: "Failed to get the prompts!"
        })
      }
    }
    getPrompts()
  }, [searchQuery, selectedTags, selectedView, page, page_limit]);

  const handleTagSelect = (tag: string) => {
    if (tag === "all") {
      setSelectedTags([]);
      return;
    }

    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter(t => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
  };

  const handleViewSelect = (view: string) => {
    setSelectedView(view);
  };

  const handleSearchChange = (query: string) => {
    setSearchQuery(query);
  };

  const handlePageLimitChange = (limit: string) => {
    setPageLimit(limit);
    setPage(1); // Reset to the first page when the page limit changes
  };

  const startItem = (page - 1) * Number(page_limit) + 1;
  const endItem = Math.min(page * Number(page_limit), totalCount);

  return (
    <div className="min-h-screen relative overflow-x-hidden">
      <AnimatedBackground />

      <div className="container mx-auto max-w-6xl">
        <Header />
        <Hero />

      </div>

      <main className="px-6 pb-20">

        <TagFilter
          selectedTags={selectedTags}
          selectedView={selectedView}
          searchQuery={searchQuery}
          onSearchChange={handleSearchChange}
          onTagSelect={handleTagSelect}
          onViewSelect={handleViewSelect}
        />


        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6 mt-4">
          {filteredPrompts.map(prompt => (
            <EnhancedPromptCard
              key={prompt.id}
              id={prompt.id}
              title={prompt.title}
              description={prompt.description}
              imageUrl={prompt.image_url}
              tags={prompt.tags}
              trending={prompt.trending}
              like_count={prompt.like_count}
            />
          ))}
        </div>

        {filteredPrompts.length === 0 && (
          <div className="text-center py-20">
            <h3 className="text-2xl text-primary mb-2">No prompts found</h3>
            <p className="text-muted-foreground">Try selecting different tags or clearing your filters.</p>
          </div>
        )}
        <div className="flex flex-col justify-center mt-10">
          <div className="flex justify-between items-center mb-4">
            <div className="flex items-center gap-2 text-sm text-white">
              <span>Showing</span>
              <strong>
          {startItem}-{endItem}
              </strong>
              <span>of</span>
              <strong>{totalCount}</strong>
              <span>items</span>
            </div>
            <Pagination className="flex-1 flex justify-center">
              <PaginationContent className="flex-wrap">
          {page > 1 && (
            <PaginationItem>
              <PaginationPrevious
                href="#"
                onClick={(e) => {
            e.preventDefault();
            setPage((prev) => Math.max(prev - 1, 1));
                }}
              />
            </PaginationItem>
          )}

          {Array.from({ length: pageCount }, (_, i) => i + 1).map((p) => {
            if (pageCount <= 5 || p === 1 || p === pageCount || (p >= page - 1 && p <= page + 1)) {
              return (
                <PaginationItem key={p}>
            <PaginationLink
              href="#"
              onClick={(e) => {
                e.preventDefault();
                setPage(p);
              }}
              isActive={p === page}
              className={p === page ? "bg-primary text-primary-foreground" : ""}
            >
              {p}
            </PaginationLink>
                </PaginationItem>
              );
            }

            if (p === 2 || p === pageCount - 1) {
              return (
                <PaginationItem key={p}>
            <PaginationEllipsis />
                </PaginationItem>
              );
            }

            return null;
          })}

          {page !== pageCount && (
            <PaginationItem>
              <PaginationNext
                href="#"
                onClick={(e) => {
            e.preventDefault();
            setPage((prev) => Math.min(prev + 1, pageCount));
                }}
              />
            </PaginationItem>
          )}
              </PaginationContent>
            </Pagination>
            <div className="flex items-center gap-2"></div>
              <span className="text-sm text-white whitespace-nowrap mr-2">Items per page</span>
              <Select
          value={page_limit}
          onValueChange={(value) => {
            handlePageLimitChange(value);
            setPage(1); // Reset to first page when changing limit
          }}
              >
          <SelectTrigger className="h-8 w-[70px]">
            <SelectValue placeholder="10" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="4">4</SelectItem>
            <SelectItem value="5">5</SelectItem>
            <SelectItem value="10">10</SelectItem>
            <SelectItem value="20">20</SelectItem>
            <SelectItem value="30">30</SelectItem>
            <SelectItem value="50">50</SelectItem>
          </SelectContent>
              </Select>
            </div>
          
        </div>
        
      </main>

      <footer className="py-8 text-center text-sm text-white">
        <p>
          © 2024 Made with ❤️ by <a href="https://github.com/Ajinkya2356" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline">Ajinkya Jagtap</a>.
        </p>
        <div className="flex justify-center space-x-4 mt-2">
          <a href="https://github.com/Ajinkya2356" target="_blank" rel="noopener noreferrer" className="hover:text-primary">

            <icons.Github className="text-white" />

          </a>
          <a href="https://linkedin.com/in/ajinkya-ai" target="_blank" rel="noopener noreferrer" className="hover:text-primary">
            <icons.Linkedin className="text-white" />
          </a>
        </div>
      </footer>
    </div>

  );
};

export default Index;
