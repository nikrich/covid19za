# Coronavirus COVID-19 (2019-nCoV) Data Repository for South Africa

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3732419.svg)](https://doi.org/10.5281/zenodo.3732419) [![arxiv](https://img.shields.io/badge/cs.CY-arXiv%3A2004.04813-B31B1B.svg)](https://arxiv.org/abs/2004.04813)

COVID 19 Data for South Africa created, maintained and hosted by [Data Science for Social Impact research group](https://dsfsi.github.io/), led by Dr. Vukosi Marivate, at the University of Pretoria. 

**Disclaimer:** We have worked to keep the data as accurate as possible. We collate the COVID 19 reporting data from NICD and DoH. We only update that data once there is an official report or statement. For the other data, we work to keep the data as accurate as possible. If you find errors. Make a pull request.

*If you use this repo for any research/development/innovation, please contact us (see contacts below)*

See our blog posts: 
* [Why we built this and how we are working](https://dsfsi.github.io/blog/covid19za-dashboard/),
* [How this is a call to action across the African continent](https://dsfsi.github.io/blog/covida19africa-call-to-action/)

*If you are interested in the **Africa-wide effort**:* Go to [https://github.com/dsfsi/covid19africa](https://github.com/dsfsi/covid19africa)

For information on daily updates on the repo, go to [https://twitter.com/vukosi/status/1239184086633242630?s=20](https://twitter.com/vukosi/status/1239184086633242630?s=20)

## Licenses

Code [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  | Data [![License: CC BY-SA 4.0](https://img.shields.io/badge/License-CC%20BY--SA%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by-sa/4.0/)

## Data Available [[/data](/data)]

**NOTE:** Since around 24 March, we have not gotten individual case data from DoH or NICD. For now if you need provincial counts use the *provincial_cumulative_timeline*. For indivividual cases up to 25 March, use the *confirmed_cases*.

| dataset         | url | raw_url[file] |
|-----------------|-----|---------------|
| provincial_cumulative_timeline|  [provincial_cumulative_timeline](/data/covid19za_provincial_cumulative_timeline_confirmed.csv)   |       [provincial_cumulative_timeline.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/covid19za_provincial_cumulative_timeline_confirmed.csv)         |
| confirmed_cases* [updated to 25 March] |  [covid19za_timeline_confirmed](/data/covid19za_timeline_confirmed.csv)   |       [covid19za_timeline_confirmed.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/covid19za_timeline_confirmed.csv)         |
| deaths |  [covid19za_timeline_deaths](/data/covid19za_timeline_deaths.csv)   |       [covid19za_timeline_deaths.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/covid19za_timeline_deaths.csv)         |
| district_data |  [district_data](/data/district_data/)   |            |
| transmission_type |  [covid19za_timeline_transmission_type](/data/covid19za_timeline_transmission_type.csv)   |       [covid19za_timeline_transmission_type.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/covid19za_timeline_transmission_type.csv)         |
| testing |  [covid19za_timeline_testing](/data/covid19za_timeline_testing.csv)   |       [covid19za_timeline_testing.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/covid19za_timeline_testing.csv)         |
|   DoH PDFs and Extracted CSVs |  [doh_pdf](/data/doh_pdf)   |              |
|   DoH Whatsapp case update archive |  [doh_whatsapp](/data/doh_whatsapp)   |              |
|   public_hospitals [validation in progress] |  [health_system_za_public_hospitals](/data/health_system_za_public_hospitals.csv)   |         [health_system_za_public_hospitals.csv](https://raw.githubusercontent.com/dsfsi/covid19za/master/data/health_system_za_public_hospitals.csv)       |

\* NICD no longer gives individual case data. Please use **provincial_cumulative_timeline** from 26 March onwards.

## Visualisation
* Google Data Studio Dashboard [URL link](https://datastudio.google.com/reporting/1b60bdc7-bec7-44c9-ba29-be0e043d8534)
![Dashboard](/visualisation/dashboard.png)
* Coronavirus Map [[Website](https://coronamap.co.za)] [[GitHub Repo](https://github.com/JayWelsh/coronamap)]
![Dashboard](/visualisation/coronamap.png)
## Data Sources:
* NICD - South Africa [URL](http://www.nicd.ac.za/media/alerts/)
* Department of Health - South Africa [Main Site](http://www.health.gov.za/), [Twitter](https://twitter.com/HealthZA/)
* South African Government Media Statements [URL](https://www.gov.za/media-statements)
* National Department of Health Data Dictionary [URL](https://dd.dhmis.org/)
* MedPages [URL](https://www.medpages.info/sf/index.php?page=homepage)
* Statistics South Africa [URL](http://www.statssa.gov.za/)

## Contributing
### Options
* *I want to help, but dont have an idea:* You can take a look at the issues to see which one you might be interested in tackling.
* *I have an idea or new feature:* Create a new issue first, assign it to yourself and then fork the repo. 
### Submitting Changes [Pull Request]
* See https://opensource.com/article/19/7/create-pull-request-github
### Resources [Get some ideas]
* [Data Science Africa COVID-19 Response](https://www.youtube.com/watch?v=9o0sa7gypMc)
* [IndabaX South Africa: Vukosi Marivate - Using data science to inform the COVID-19 outbreak in Africa](https://www.youtube.com/watch?v=DZOpypSA85I)
* [Stanford <> CS472 Data science and AI for COVID-19](https://sites.google.com/view/data-science-covid-19)
## Contributors
<a href="https://github.com/dsfsi/covid19za/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=dsfsi/covid19za" />
</a>
Made with [contributors-img](https://contributors-img.web.app).
* See https://github.com/dsfsi/covid19za/graphs/contributors

## Contact
* Vukosi Marivate - vukosi.marivate@cs.up.ac.za, [@vukosi](https://twitter.com/vukosi)

## Citing the dataset
**On a visualisation/notebook/webapp:**

> Data Science for Social Impact Research Group @ University of Pretoria, *Coronavirus COVID-19 (2019-nCoV) Data Repository for South Africa.* Available on: https://github.com/dsfsi/covid19za.


**In a publication**

Arxiv Preprint
> @misc{marivate2020framework,
    title={A Framework For Sharing Publicly Available Data To Inform The COVID-19 Outbreak in Africa: A South African Case Study},
    author={Vukosi Marivate and Herkulaas MvE Combrink},
    year={2020},
    eprint={2004.04813},
    archivePrefix={arXiv},
    primaryClass={cs.CY},
    url = {[https://arxiv.org/abs/2004.04813](https://arxiv.org/abs/2004.04813)}
}

and Dataset
> @dataset{marivate_vukosi_2020_3732419,
  author       = {Marivate, Vukosi and
                  de Waal, Alta and
                  Combrink, Herkulaas and
                  Lebogo, Ofentswe and
                  Moodley, Shivan and
                  Mtsweni, Nompumelelo and
                  Rikhotso, Vuthlari and
                  Welsh, Jay and
                  Mkhondwane, S'busiso},
  title        = {{Coronavirus disease (COVID-19) case data - South 
                   Africa}},
  month        = mar,
  year         = 2020,
  publisher    = {Zenodo},
  doi          = {10.5281/zenodo.3732419},
  url          = {[https://doi.org/10.5281/zenodo.3732419](https://doi.org/10.5281/zenodo.3732419)}
}

### Showcase

* [Covid19 SA Data](https://simonrosen173.github.io/Covid19SAData/)
