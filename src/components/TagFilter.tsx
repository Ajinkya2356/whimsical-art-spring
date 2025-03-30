
import React from 'react';
import { Button } from '@/components/ui/button';
import { LayoutGrid, Tag, TrendingUp, Heart } from 'lucide-react';
import { ScrollArea } from '@/components/ui/scroll-area';

const tags = [
  { id: "all", name: "All Prompts", icon: <LayoutGrid size={18} /> },
  { id: "Forest", name: "Forest", icon: <Tag size={18} /> },
  { id: "Spirit", name: "Spirit", icon: <Tag size={18} /> },
  { id: "Landscape", name: "Landscape", icon: <Tag size={18} /> },
  { id: "Character", name: "Character", icon: <Tag size={18} /> },
  { id: "Magic", name: "Magic", icon: <Tag size={18} /> },
  { id: "Sky", name: "Sky", icon: <Tag size={18} /> },
  { id: "Ocean", name: "Ocean", icon: <Tag size={18} /> },
  { id: "Castle", name: "Castle", icon: <Tag size={18} /> },
  { id: "Animal", name: "Animal", icon: <Tag size={18} /> },
  { id: "Fantasy", name: "Fantasy", icon: <Tag size={18} /> }
];

const views = [
  { id: "all", name: "All", icon: <LayoutGrid size={18} /> },
  { id: "trending", name: "Trending", icon: <TrendingUp size={18} /> },
  { id: "favorites", name: "Favorites", icon: <Heart size={18} /> },
];

interface TagFilterProps {
  selectedTags: string[];
  selectedView: string;
  onTagSelect: (tag: string) => void;
  onViewSelect: (view: string) => void;
}

const TagFilter = ({ selectedTags, selectedView, onTagSelect, onViewSelect }: TagFilterProps) => {
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
              className={`flex items-center gap-2 ${
                selectedView === view.id 
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
      
      <ScrollArea className="w-full pb-4">
        <div className="flex flex-nowrap gap-2 py-2 px-1 overflow-x-auto hide-scrollbar">
          {tags.map(tag => (
            <Button 
              key={tag.id}
              variant={(tag.id === "all" && selectedTags.length === 0) || selectedTags.includes(tag.id) ? "default" : "outline"}
              size="sm"
              className="flex items-center gap-2 rounded-full transition-all whitespace-nowrap flex-shrink-0"
              onClick={() => onTagSelect(tag.id)}
            >
              {tag.icon}
              <span>{tag.name}</span>
            </Button>
          ))}
        </div>
      </ScrollArea>
      
      <style jsx>{`
        .hide-scrollbar::-webkit-scrollbar {
          display: none;
        }
        .hide-scrollbar {
          -ms-overflow-style: none;
          scrollbar-width: none;
        }
      `}</style>
    </div>
  );
};

export default TagFilter;
