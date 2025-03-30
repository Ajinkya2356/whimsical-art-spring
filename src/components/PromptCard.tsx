
import React from 'react';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Eye } from 'lucide-react';

export interface PromptCardProps {
  id: string;
  title: string;
  description: string;
  imageUrl: string;
  tags: string[];
}

const PromptCard = ({ title, description, imageUrl, tags }: PromptCardProps) => {
  return (
    <div className="ghibli-card group animate-fadeIn" style={{ animationDelay: Math.random() * 0.5 + 's' }}>
      <div className="relative h-48 overflow-hidden">
        <img 
          src={imageUrl} 
          alt={title}
          className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110"
        />
        <div className="absolute inset-0 bg-gradient-to-t from-black/40 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300">
          <div className="absolute bottom-3 right-3">
            <Button size="sm" variant="secondary" className="rounded-full bg-white/80 backdrop-blur-sm hover:bg-white">
              <Eye className="h-4 w-4 mr-1" /> View
            </Button>
          </div>
        </div>
      </div>
      <div className="p-4">
        <h3 className="text-xl text-ghibli-forest">{title}</h3>
        <p className="text-sm text-muted-foreground mt-2 line-clamp-2">{description}</p>
        <div className="flex flex-wrap gap-2 mt-3">
          {tags.map((tag) => (
            <Badge key={tag} variant="outline" className="bg-ghibli-sky/10 text-ghibli-forest border-ghibli-sky/20">
              {tag}
            </Badge>
          ))}
        </div>
      </div>
    </div>
  );
};

export default PromptCard;
