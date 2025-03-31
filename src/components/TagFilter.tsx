import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { LayoutGrid, Tag, TrendingUp, Heart } from 'lucide-react';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Input } from '@/components/ui/input';
import { Search } from 'lucide-react';
import supabase from '../utils/supabase';
import { useToast } from '@/hooks/use-toast';

const views = [
  { id: "all", name: "All", icon: <LayoutGrid size={18} /> },
  { id: "trending", name: "Trending", icon: <TrendingUp size={18} /> },
];

interface TagFilterProps {
  selectedTags: string[];
  selectedView: string;
  searchQuery: string;
  onTagSelect: (tag: string) => void;
  onViewSelect: (view: string) => void;
  onSearchChange: (query: string) => void;
}

const TagFilter = ({ selectedTags, selectedView, searchQuery, onTagSelect, onViewSelect, onSearchChange }: TagFilterProps) => {
  const [uniqueTags, setUniqueTags] = useState<string[]>([]);
  const { toast } = useToast();

  useEffect(() => {
    async function fetchUniqueTags() {
      try {
        const { data, error } = await supabase.rpc('get_unique_tags', {});

        if (error) {
          console.error("Error fetching unique tags:", error);
          toast({
            title: "Error",
            description: "Failed to fetch unique tags. Please try again.",
          });
        } else {
          // Assuming data is an array of objects like { tag: "tagName" }
          setUniqueTags(data?.map(item => item.tag) || []);
        }
      } catch (error) {
        console.error("Error fetching unique tags:", error);
        toast({
          title: "Error",
          description: "Failed to fetch unique tags. Please try again.",
        });
      }
    }

    fetchUniqueTags();
  }, []);

  return (
    <div className="w-full py-6 space-y-4">
      <div className="flex items-center justify-between mb-2">
        <h2 className="text-xl font-semibold text-primary">Browse Prompts</h2>
        <div className="flex gap-1 bg-muted/30 rounded-lg p-1">
          {views.map(view => (
            <Button
              key={view.id}
              variant="ghost"
              size="sm"
              className={`flex items-center gap-2 ${selectedView === view.id
                ? "bg-card text-primary"
                : "text-muted-foreground hover:text-primary"
                }`}
              onClick={() => onViewSelect(view.id)}
            >
              {view.icon}
              <span className="hidden sm:inline">{view.name}</span>
            </Button>
          ))}
        </div>
      </div>
      <div className="w-full max-w-xl mx-auto relative animate-fadeIn" style={{ animationDelay: '0.3s' }}>
        <div className="relative">
          <Search className="absolute left-3 top-3 h-5 w-5 text-muted-foreground/70" />
          <Input
            placeholder="Search for magical prompts..."
            className="pl-10 py-6 bg-muted/50 border-muted text-primary focus-visible:ring-ghibli-accent w-full rounded-full"
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
          />
        </div>
      </div>
      <ScrollArea className="w-full">
        <div className="flex flex-wrap gap-2 py-2 px-1">
          <Button
            key="all"
            variant={selectedTags.length === 0 ? "default" : "outline"}
            size="sm"
            className={`flex items-center gap-2 rounded-full transition-all`}
            onClick={() => onTagSelect("all")}
          >
            <LayoutGrid size={18} />
            <span>All Prompts</span>
          </Button>
          {uniqueTags.map(tag => (
            <Button
              key={tag}
              variant={selectedTags.includes(tag) ? "default" : "outline"}
              size="sm"
              className={`flex items-center gap-2 rounded-full transition-all`}
              onClick={() => onTagSelect(tag)}
            >
              <Tag size={18} />
              <span>{tag}</span>
            </Button>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
};

export default TagFilter;
