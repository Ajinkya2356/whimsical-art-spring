
import React, { useState } from 'react';
import Header from '@/components/Header';
import Hero from '@/components/Hero';
import AnimatedBackground from '@/components/AnimatedBackground';
import PromptCard, { PromptCardProps } from '@/components/PromptCard';
import TagFilter from '@/components/TagFilter';

// Mock data - in a real app, this would come from Supabase
const mockPrompts: PromptCardProps[] = [
  {
    id: "1",
    title: "Spirit of the Forest",
    description: "A gentle forest spirit with glowing eyes, surrounded by tiny floating light orbs in a misty ancient forest at dusk.",
    imageUrl: "https://images.unsplash.com/photo-1518495973542-4542c06a5843",
    tags: ["Forest", "Spirit", "Magic"]
  },
  {
    id: "2",
    title: "Sky Castle Journey",
    description: "A flying castle drifting among sunset clouds, with hanging gardens and waterfalls cascading from its edges.",
    imageUrl: "https://images.unsplash.com/photo-1482938289607-e9573fc25ebb",
    tags: ["Sky", "Castle", "Fantasy"]
  },
  {
    id: "3",
    title: "River Guardian",
    description: "A wise dragon spirit coiled in a crystal-clear river, scales shimmer like opals under the moonlight.",
    imageUrl: "https://images.unsplash.com/photo-1472396961693-142e6e269027",
    tags: ["Animal", "Spirit", "Magic", "Fantasy"]
  },
  {
    id: "4",
    title: "Ancient Bathhouse",
    description: "A traditional Japanese bathhouse perched on a cliff, steam rising, lanterns glowing, with spirit creatures as guests.",
    imageUrl: "https://images.unsplash.com/photo-1470071459604-3b5ec3a7fe05",
    tags: ["Spirit", "Fantasy", "Character"]
  },
  {
    id: "5",
    title: "Wind Valley",
    description: "Rolling green hills under a vast blue sky, with wildflowers dancing in the wind and small cottages nestled in the valleys.",
    imageUrl: "https://images.unsplash.com/photo-1509316975850-ff9c5deb0cd9",
    tags: ["Landscape", "Sky"]
  },
  {
    id: "6",
    title: "Whisper of the Waves",
    description: "A small fishing village on stilts above a tranquil sea, with paper lanterns reflecting in the water at twilight.",
    imageUrl: "https://images.unsplash.com/photo-1500673922987-e212871fec22",
    tags: ["Ocean", "Landscape"]
  }
];

const Index = () => {
  const [selectedTags, setSelectedTags] = useState<string[]>([]);

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

  const filteredPrompts = selectedTags.length > 0
    ? mockPrompts.filter(prompt => 
        selectedTags.some(tag => prompt.tags.includes(tag))
      )
    : mockPrompts;

  return (
    <div className="min-h-screen relative overflow-x-hidden">
      <AnimatedBackground />
      
      <div className="container mx-auto max-w-7xl">
        <Header />
        <Hero />
        
        <main className="px-6 pb-20">
          <TagFilter selectedTags={selectedTags} onTagSelect={handleTagSelect} />
          
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mt-4">
            {filteredPrompts.map(prompt => (
              <PromptCard key={prompt.id} {...prompt} />
            ))}
          </div>
          
          {filteredPrompts.length === 0 && (
            <div className="text-center py-20">
              <h3 className="text-2xl text-ghibli-forest mb-2">No prompts found</h3>
              <p className="text-muted-foreground">Try selecting different tags or clearing your filters.</p>
            </div>
          )}
        </main>
        
        <footer className="py-8 text-center text-sm text-muted-foreground">
          <p>© 2023 Ghibli Prompt Garden. Inspired by the magical worlds of Studio Ghibli.</p>
        </footer>
      </div>
    </div>
  );
};

export default Index;
