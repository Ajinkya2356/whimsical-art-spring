
import React from 'react';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';

const tags = [
  "Forest", "Spirit", "Landscape", "Character", "Magic", 
  "Sky", "Ocean", "Castle", "Animal", "Fantasy"
];

interface TagFilterProps {
  selectedTags: string[];
  onTagSelect: (tag: string) => void;
}

const TagFilter = ({ selectedTags, onTagSelect }: TagFilterProps) => {
  return (
    <div className="w-full py-6 overflow-x-auto scrollbar-thin">
      <div className="flex gap-2 pb-2 min-w-max">
        <Button 
          variant={selectedTags.length === 0 ? "default" : "outline"} 
          size="sm"
          className={selectedTags.length === 0 ? "bg-ghibli-forest text-white" : "text-ghibli-forest border-ghibli-forest/30"}
          onClick={() => onTagSelect("all")}
        >
          All Prompts
        </Button>
        {tags.map(tag => (
          <Badge 
            key={tag}
            variant="outline"
            className={`text-sm cursor-pointer px-3 py-1 ${
              selectedTags.includes(tag) 
                ? "bg-ghibli-forest text-white" 
                : "bg-ghibli-sky/10 text-ghibli-forest border-ghibli-sky/20 hover:bg-ghibli-sky/20"
            }`}
            onClick={() => onTagSelect(tag)}
          >
            {tag}
          </Badge>
        ))}
      </div>
    </div>
  );
};

export default TagFilter;
