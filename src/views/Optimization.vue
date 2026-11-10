
<template>
  <div class="optimization-page">
    <!-- HERO -->
    <section class="hero">
      <div class="hero-grid"></div>

      <div class="hero-content">
        <div class="eyebrow">
          <span class="pulse"></span>
          NJIT ACM PRESENTS
          <span class="year">2026</span>
        </div>

        <h1>
          OPTIMI<span class="accent">Z</span>ATION
          <span class="hero-subtitle">PUSHING TECHNOLOGY TO ITS LIMITS.</span>
        </h1>

        <p class="hero-description">
          A new computing competition!
          Optimize! Optimize! Optimize!
        </p>

        <div class="hero-actions">
          <a href="#challenge" class="pill primary-btn">
            Explore challenge
          </a>
          <a href="#about" class="pill primary-btn">
            Discover event
          </a>
        </div>

        <div class="hero-meta">
          <span>Algorithms</span>
          <span>Engineering</span>
          <span>Optimization</span>
        </div>
      </div>

      <div class="hero-visual">
        <div class="visual-label">System Efficiency</div>
        <div class="orbit orbit-one"></div>
        <div class="orbit orbit-two"></div>
        <div class="orbit orbit-three"></div>
        <div class="visual-core">
          <span>OPT</span>
          <strong>∞</strong>
        </div>
        <div class="visual-coordinate">
          40.7421° N<br />
          74.1784° W
        </div>
      </div>
    </section>

    <!-- EVENT INTRO -->
    <section id="about" class="section intro">
      <div class="section-index">Main Idea</div>

      <div>
        <h2>
          Focusing on expansion<br />
          <span class="muted">Moving from building to fine-tuning</span>
        </h2>

        <p>
          This event is a hands-on computing competition
          where participants take an existing process and find ways
          to make it more efficient. The goal is to get the most out
          of the available resources while supporting as many users
          or requests as possible.
        </p>
      </div>
    </section>

    <!-- CHALLENGE -->
    <section id="challenge" class="section challenge">
      <div class="section-index">Challenge</div>

      <h2>
        To Find the<br />
        BEST<br />
        Solution
      </h2>

      <div class="challenge-grid">
        <article class="challenge-card">
          <span class="number">01</span>
          <div class="card-icon">⌘</div>
          <h3>ALGORITHMS</h3>
          <p>
            Design an efficient system able to handle
            many requests without breaking down.
          </p>
          <span class="tag">Computational Thinking</span>
        </article>

        <article class="challenge-card">
          <span class="number">02</span>
          <div class="card-icon">◈</div>
          <h3>System Design</h3>
          <p>
            Translate requirements into scalable architectures
            while balancing performance and constraints.
          </p>
          <span class="tag">ENGINEERING DESIGN</span>
        </article>

        <article class="challenge-card">
          <span class="number">03</span>
          <div class="card-icon">⌬</div>
          <h3>IMPLEMENTATION</h3>
          <p>
            Build functional software, test assumptions,
            then iterate toward a stronger solution.
          </p>
          <span class="tag">SOFTWARE DEVELOPMENT</span>
        </article>

        <article class="challenge-card">
          <span class="number">04</span>
          <div class="card-icon">↗</div>
          <h3>PERFORMANCE</h3>
          <p>
            Compare solutions. Measure where tradeoffs matter.
            Make a better solution.
          </p>
          <span class="tag">TESTING &amp; VALIDATION</span>
        </article>
      </div>
    </section>

    <!-- CAPSTONE -->
    <section class="capstone">
      <div class="capstone-content">
        <div class="section-index">ENGINEERING FUNDAMENTALS</div>

        <h2>
          FROM REQUIREMENTS<br />
          TO <span>IMPLEMENTATION.</span>
        </h2>

        <p>
          Based on the fundamentals taught in YWCC Capstone,
          participants will work through the different stages
          of an engineering project. This includes understanding
          the problem, planning a design, developing a solution,
          and testing the final result.
        </p>

        <div class="process">
          <div><b>01</b><span>REQUIREMENTS</span></div>
          <div><b>02</b><span>DESIGN</span></div>
          <div><b>03</b><span>DEVELOP</span></div>
          <div><b>04</b><span>TEST</span></div>
          <div><b>05</b><span>PRESENT</span></div>
        </div>
      </div>
    </section>

    <!-- FORMAT -->
    <section class="section format">
      <div class="section-index">HOW IT WORKS</div>

      <h2>
        HOW THE<br />
        EVENT WORKS.
      </h2>

      <div class="timeline">
        <div class="timeline-item">
          <span>01</span>
          <div>
            <h3>PROBLEM BRIEFING</h3>
            <p>Review the problem, objectives, and available resources.</p>
          </div>
        </div>

        <div class="timeline-item">
          <span>02</span>
          <div>
            <h3>PLAN &amp; DESIGN</h3>
            <p>Decide how to approach the problem and what needs to be improved.</p>
          </div>
        </div>

        <div class="timeline-item">
          <span>03</span>
          <div>
            <h3>BUILD &amp; IMPROVE</h3>
            <p>Develop a solution, test different approaches, and make changes.</p>
          </div>
        </div>

        <div class="timeline-item">
          <span>04</span>
          <div>
            <h3>TEST &amp; COMPARE</h3>
            <p>Measure performance and compare results against the original process.</p>
          </div>
        </div>

        <div class="timeline-item">
          <span>05</span>
          <div>
            <h3>FINAL PRESENTATION</h3>
            <p>Show the final solution and explain what improved and why.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="closing">
      <div class="section-index">THE NEXT CHALLENGE</div>

      <h2>
        BUILD IT.<br />
        <span>MAKE IT BETTER.</span>
      </h2>

      <p>Take an existing process and find a better way to run it.</p>

      <a href="#about" class="pill primary-btn">
        EXPLORE OPTIMIZATION ↑
      </a>

      <div class="closing-footer">
        <span>NJIT ACM · OPTIMIZATION 2026</span>
        <span>MOORE'S LAW IS GONE.</span>
      </div>
    </section>
  </div>
</template>

<style scoped>

/* PAGE */
.optimization-page {
  width: 100%;
  overflow: hidden;
  color: var(--text-color);
}

/* HERO */
.hero {
  position: relative;
  min-height: 85vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  isolation: isolate;
  padding: 60px 20px;
}

.hero-grid {
  position: absolute;
  inset: 0;
  opacity: 0.12;
  background-image:
    linear-gradient(#aaa 1px, transparent 1px),
    linear-gradient(90deg, #aaa 1px, transparent 1px);
  background-size: 55px 55px;
  pointer-events: none;
}

.hero-content {
  position: relative;
  z-index: 2;
  max-width: 850px;
}

.eyebrow {
  color: var(--hacknjit-tertiary);
  letter-spacing: 0.2em;
  font-size: 0.85rem;
  margin-bottom: 25px;
}

.pulse {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--hacknjit-tertiary);
  margin-right: 10px;
  box-shadow: 0 0 12px var(--hacknjit-tertiary);
}

.year {
  color: #aaa;
  margin-left: 10px;
}

.hero h1 {
  font-size: clamp(3rem, 9vw, 6em);
  line-height: 1;
  letter-spacing: -0.04em;
}

.accent,
.capstone h2 span,
.closing h2 span {
  color: var(--hacknjit-tertiary);
}

.hero-subtitle {
  display: block;
  font-size: 0.25em;
  letter-spacing: 0.2em;
  color: #aaa;
  margin-top: 20px;
}

.hero-description {
  margin: 25px auto;
  color: #ddd;
}

.hero-actions {
  display: flex;
  justify-content: center;
  gap: 14px;
  flex-wrap: wrap;
  margin: 30px 0;
}

.primary-btn,
.secondary-btn {
  display: inline-block;
  text-decoration: none;
  font-size: 1rem;
  font-weight: bold;
  transition: transform 250ms ease, filter 300ms ease;
}

.secondary-btn {
  padding: 8px 24px;
  border: 1px solid var(--hacknjit-secondary);
  border-radius: 1000px;
  background: var(--hacknjit-primary-transparent);
  color: white;
}

.secondary-btn:hover {
  transform: scale(var(--hover-scale));
  filter: drop-shadow(0 8px 10px #00000055);
}

.hero-meta {
  display: flex;
  justify-content: center;
  gap: 25px;
  color: #aaa;
  font-size: 0.85rem;
  letter-spacing: 0.12em;
}

/* HERO ORBIT */
.hero-visual {
  position: absolute;
  top: 18%;
  right: 8%;
  width: 230px;
  height: 230px;
  z-index: 0;
  opacity: 0.3;
  pointer-events: none;
}

.visual-label,
.visual-coordinate {
  color: var(--hacknjit-tertiary);
  font-size: 0.7rem;
  letter-spacing: 0.15em;
}

.orbit {
  position: absolute;
  inset: 25px;
  border: 1px solid var(--hacknjit-tertiary);
  border-radius: 50%;
  transform: rotate(-25deg) scaleY(0.45);
}

.orbit-two {
  transform: rotate(45deg) scaleY(0.45);
}

.orbit-three {
  transform: rotate(90deg) scaleY(0.45);
}

.visual-core {
  position: absolute;
  inset: 75px;
  border: 1px solid var(--hacknjit-tertiary);
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: var(--hacknjit-tertiary);
}

.visual-core strong {
  font-size: 35px;
}

.visual-coordinate {
  position: absolute;
  bottom: 0;
  right: 0;
}

/* INTRO */
.intro {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 32px;
  align-items: start;
  padding-top: 65px;
  padding-bottom: 65px;
}

.intro h2 {
  text-align: left;
  margin-bottom: 25px;
}

.muted {
  color: #aaa;
  font-size: 0.55em;
}

.intro p {
  max-width: 720px;
  color: #ddd;
}

/* CHALLENGE */
.challenge {
  padding-top: 45px;
  padding-bottom: 70px;
}

.challenge > h2 {
  margin: 10px 0 40px;
}

.challenge-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.challenge-card {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  background: var(--hacknjit-primary-transparent);
  border: 1px solid var(--edge-colors);
  border-radius: var(--border-radius);
  padding: 24px;
  min-height: 270px;
}

.number {
  color: var(--hacknjit-tertiary);
}

.card-icon {
  font-size: 2rem;
  color: var(--hacknjit-tertiary);
  margin: 15px 0;
}

.challenge-card h3 {
  text-align: left;
  font-size: 1.2rem;
  margin-bottom: 12px;
}

.challenge-card p {
  font-size: 0.95rem;
  color: #ddd;
}

.tag {
  color: var(--hacknjit-tertiary);
  font-size: 0.7rem;
  margin-top: auto;
  padding-top: 20px;
  letter-spacing: 0.08em;
}

/* CAPSTONE */
.capstone {
  max-width: var(--max-content-width);
  width: calc(100% - 40px);
  margin: 50px auto;
  background: var(--hacknjit-primary-transparent);
  border: 1px solid var(--edge-colors);
  border-radius: var(--border-radius);
  padding: 55px 6%;
}

.capstone h2 {
  margin: 15px 0 20px;
}

.capstone p {
  max-width: 720px;
  margin: auto;
  text-align: center;
  color: #ddd;
}

.process {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 12px;
  margin-top: 40px;
}

.process div {
  display: flex;
  flex-direction: column;
  gap: 8px;
  border-top: 1px solid var(--hacknjit-tertiary);
  padding-top: 12px;
}

.process b {
  color: var(--hacknjit-tertiary);
}

.process span {
  font-size: 0.75rem;
  letter-spacing: 0.08em;
}

/* FORMAT */
.format {
  padding-top: 70px;
  padding-bottom: 70px;
}

.format > h2 {
  margin: 15px 0 35px;
}

.timeline {
  max-width: 850px;
  margin: 0 auto;
}

.timeline-item {
  display: grid;
  grid-template-columns: 60px 1fr;
  gap: 20px;
  padding: 22px 0;
  border-bottom: 1px solid #ffffff35;
}

.timeline-item > span {
  color: var(--hacknjit-tertiary);
  font-size: 1.2rem;
}

.timeline-item h3 {
  text-align: left;
  font-size: 1.25rem;
  margin-bottom: 5px;
}

.timeline-item p {
  color: #ddd;
  font-size: 0.95rem;
}

/* CLOSING */
.closing {
  text-align: center;
  padding: 100px 24px 35px;
}

.closing h2 {
  margin: 20px 0;
}

.closing p {
  margin-bottom: 25px;
  color: #ddd;
}

.closing-footer {
  display: flex;
  justify-content: space-between;
  gap: 15px;
  flex-wrap: wrap;
  border-top: 1px solid #ffffff44;
  margin-top: 90px;
  padding-top: 20px;
  font-size: 0.8rem;
  letter-spacing: 0.12em;
}

/* RESPONSIVE */
@media (max-width: 900px) {
  .challenge-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .process {
    grid-template-columns: repeat(2, 1fr);
  }

  .hero-visual {
    opacity: 0.15;
  }
}

@media (max-width: 550px) {
  .hero {
    min-height: 75vh;
  }

  .hero-meta {
    gap: 10px;
    font-size: 0.7rem;
  }

  .hero-visual {
    right: 0;
  }

  .intro {
    grid-template-columns: 1fr;
  }

  .challenge-grid,
  .process {
    grid-template-columns: 1fr;
  }

  .challenge-card {
    min-height: 0;
  }

  .capstone {
    padding: 24px;
  }

  .timeline-item {
    grid-template-columns: 40px 1fr;
    gap: 12px;
  }

  .closing-footer {
    justify-content: center;
    text-align: center;
  }
}
</style>