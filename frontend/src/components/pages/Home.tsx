import { useState, useEffect } from 'react'
import welcomeImage from '../../assets/welcome.png'
import educationImage from '../../assets/books.png'
import healthImage from '../../assets/medical.png'

const images = [
  { src: welcomeImage, alt: 'Welcome' },
  { src: educationImage, alt: 'Education' },
  { src: healthImage, alt: 'Medicine' },
  
]

export default function Home() {
  const [current, setCurrent] = useState(0)
  const user = JSON.parse(localStorage.getItem('user') || '{}')

  const prev = () => setCurrent(c => (c === 0 ? images.length - 1 : c - 1))
  const next = () => setCurrent(c => (c === images.length - 1 ? 0 : c + 1))

  useEffect(() => {
    const id = setInterval(next, 5000)
    return () => clearInterval(id)
  }, [])

  return (
    <div className="home-page">
      <h1>Welcome, {user.firstName}!</h1>

      <div className="image-carousel">
        <button onClick={prev} className="carousel-arrow">‹</button>

        <div className="carousel-viewport">
          <div
            className="carousel-track"
            style={{ transform: `translateX(-${current * 100}%)` }}
          >
            {images.map((img, i) => (
              <div className="carousel-slide" key={i}>
                <img src={img.src} alt={img.alt} className="welcome-image" />
              </div>
            ))}
          </div>
        </div>

        <button onClick={next} className="carousel-arrow">›</button>
      </div>

      <div className="carousel-dots">
        {images.map((_, i) => (
          <button
            key={i}
            className={`carousel-dot ${i === current ? 'active' : ''}`}
            onClick={() => setCurrent(i)}
            aria-label={`Go to slide ${i + 1}`}
          />
        ))}
      </div>
    </div>
  )
}