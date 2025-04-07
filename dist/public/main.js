window.addEventListener("load", () => {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        const lasers = Array.from(entry.target.querySelectorAll("[data-laser]"));

        if (entry.isIntersecting) {
          lasers.forEach((laser) => {
            laser.classList.remove("opacity-0")
          });
        }else {
          lasers.forEach((laser) => {
            laser.classList.add("opacity-0")
          });
        }
      });
    },
    {
      root: null,
      rootMargin: "-100px 0px",
      threshold: [1],
    }
  );

  if (window.innerWidth <= 768) {
    const projects = Array.from(document.querySelectorAll(".project"));
    projects.forEach((element) => {
      observer.observe(element);
    });
  }
});
