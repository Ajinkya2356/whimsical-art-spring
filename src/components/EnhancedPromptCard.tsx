
import React from 'react';
import { Badge } from '@/components/ui/badge';
import PromptActions from './PromptActions';

export interface EnhancedPromptCardProps {
  id: string;
  title: string;
  description: string;
  imageUrl: string;
  tags: string[];
  trending?: boolean;
}

const EnhancedPromptCard = ({ id, title, description, imageUrl, tags, trending = false }: EnhancedPromptCardProps) => {
  return (
    <div className="glass-card rounded-xl overflow-hidden animate-fadeIn transform transition-all duration-300 hover:scale-[1.02] hover:shadow-xl">
      {imageUrl && (
        <div className="w-full h-48 overflow-hidden">
          <img 
            src={imageUrl} 
            alt={title}
            className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110"
          />
        </div>
      )}
      
      <div className="p-5 space-y-4">
        {trending && (
          <Badge className="bg-ghibli-accent text-background mb-2">
            Trending
          </Badge>
        )}
        
        <h3 className="text-xl font-bold text-primary">{title}</h3>
        
        <div className="text-sm text-muted-foreground">
          {description}
        </div>
        
        <div className="flex flex-wrap gap-1 mt-2">
          {tags.map(tag => (
            <Badge 
              key={tag} 
              variant="outline" 
              className="text-xs bg-muted/40 text-muted-foreground border-muted/50"
            >
              {tag}
            </Badge>
          ))}
        </div>
        
        <PromptActions promptId={id} promptText={description} />
      </div>
    </div>
  );
};

export default EnhancedPromptCard;
