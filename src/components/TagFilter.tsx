
import React from 'react';
import { Button } from '@/components/ui/button';
import { Folder, Tag, List, LayoutGrid, Heart, TrendingUp, Star } from 'lucide-react';

const tags = [
  { id: "all", name: "All Prompts", icon: <LayoutGrid size={18} /> },
  { id: "forest", name: "Forest", icon: <Tag size={18} /> },
  { id: "spirit", name: "Spirit", icon: <Tag size={18} /> },
  { id: "landscape", name: "Landscape", icon: <Tag size={18} /> },
  { id: "character", name: "Character", icon: <Tag size={18} /> },
  { id: "magic", name: "Magic", icon: <Tag size={18} /> },
  { id: "sky", name: "Sky", icon: <Tag size={18} /> },
  { id: "ocean", name: "Ocean", icon: <Tag size={18} /> },
  { id: "castle", name: "Castle", icon: <Tag size={18} /> },
  { id: "animal", name: "Animal", icon: <Tag size={18} /> },
  { id: "fantasy", name: "Fantasy", icon: <Tag size={18} /> }
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
        <div className="flex gap-1 bg-muted rounded-lg p-1">
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
      
      <div className="overflow-x-auto scrollbar-thin pb-2">
        <div className="flex gap-2 min-w-max">
          {tags.map(tag => (
            <div 
              key={tag.id}
              className={`category-pill ${
                tag.id === "all" && selectedTags.length === 0 || selectedTags.includes(tag.id)
                  ? "active" 
                  : "inactive"
              }`}
              onClick={() => onTagSelect(tag.id)}
            >
              {tag.icon}
              <span>{tag.name}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default TagFilter;
