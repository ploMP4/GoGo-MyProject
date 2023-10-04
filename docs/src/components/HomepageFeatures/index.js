import React from "react";
import clsx from "clsx";
import styles from "./styles.module.css";

const FeatureList = [
  {
    title: "Easy to Use",
    Svg: require("@site/static/img/easytouse.svg").default,
    description: (
      <>
        GoGo was designed to be easily configured by making use of the tools
        you're already used to in your everyday life.
      </>
    ),
  },
  {
    title: "Focus on What Matters",
    Svg: require("@site/static/img/focus.svg").default,
    description: (
      <>
        Don't think about the repetitive actions and how to setup things.
        Instead configure the process once and reuse it whenever you need to.
      </>
    ),
  },
  {
    title: "Flexible",
    Svg: require("@site/static/img/flex.svg").default,
    description: (
      <>
        By combining popular cli tools and the ability to template and
        manipulate files you can overcome any tedious task that gets in your
        way.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center" style={{ marginTop: 20 }}>
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
