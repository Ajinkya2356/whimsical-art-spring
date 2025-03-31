import React, { useState, useEffect } from 'react';
import Header from '@/components/Header';
import Hero from '@/components/Hero';
import AnimatedBackground from '@/components/AnimatedBackground';
import TagFilter from '@/components/TagFilter';
import EnhancedPromptCard from '@/components/EnhancedPromptCard';
import { icons } from 'lucide-react';
import supabase from '../utils/supabase';
import { useToast } from '@/hooks/use-toast';


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

  // Filter prompts whenever search query, tags, or view changes
  useEffect(() => {
    async function getPrompts() {
      try {
        const { data, error } = await supabase.rpc('get_prompts', {
          page: 1,
          page_limit: 20,
          search_tags: selectedTags.length > 0 ? selectedTags : undefined,
          is_trending: selectedView === "trending" ? true : undefined,
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
        }
      } catch (e) {
        toast({
          title: 'Error',
          description: "Failed to get the prompts!"
        })
      }
    }
    getPrompts()
  }, [searchQuery, selectedTags, selectedView]);

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
              trending={selectedView == "trending"}
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
      </main>

      <footer className="py-8 text-center text-sm text-muted-foreground">
        <p>
          © 2024 Made with ❤️ by <a href="https://github.com/Ajinkya2356" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline">Ajinkya Jagtap</a>.
        </p>
        <div className="flex justify-center space-x-4 mt-2">
          <a href="https://github.com/Ajinkya2356" target="_blank" rel="noopener noreferrer" className="hover:text-primary">

            <icons.Github className="text-muted-foreground" />

          </a>
          <a href="https://linkedin.com/in/ajinkya-ai" target="_blank" rel="noopener noreferrer" className="hover:text-primary">
            <icons.Linkedin className="text-muted-foreground" />
          </a>
        </div>
      </footer>
    </div>

  );
};

export default Index;
