<!-- Registration.vue -->
<template>
  <Vueform 
    v-bind="vueform" 
    @submit="customSubmit"
  />
</template>
<script>
import { useVueform, Vueform } from '@vueform/vueform'
import schoolsJson from '../assets/schools.json'
import countriesJson from '../assets/countries.json'

export default {
  mixins: [Vueform],
  setup: useVueform,
  data: () => ({
    vueform: {
      size: 'md',
      displayErrors: true,
      endpoint: '/api/register',
      method: 'POST',
      uploadFileOnChange: false, // Disable automatic file (PDF) uploads
      enctype: 'multipart/form-data', // For PDF (resume) uploads
      validateOn: 'step|change',
      steps: {
        page0: {
          label: 'Personal Information',
          elements: [
            'HackNJIT Registration',
            'divider',
            'container',
            'preferred_name_container',
            'age',
            'phone',
            'email',
            'country',
            'uni',
            'lvlofstudy',
            'firsthack',
            'major',
            'shirtsize',
            'dietaryrestrictions',
            'divider_1',
          ],
        },
        page1: {
          label: 'MLH Agreements',
          elements: [
            'h1',
            'divider_4',
            'terms',
            'marketing_emails',
            'marketing_emails_1',
            'divider_3',
          ],
        },
        page2: {
          elements: [
            'h1_2',
            'p_1',
            'minority',
            'gender',
            'race',
            'divider_6',
            'h1_1',
            'p',
            'divider_2',
            // 'resume',
            'linkedin',
            'divider_5',
          ],
          label: 'Extra Information',
        },
      },
      addClass: 'vf-registration',
      schema: {
        'HackNJIT Registration': {
          type: 'static',
          content: 'HackNJIT Registration',
          tag: 'h1',
        },
        divider: {
          type: 'static',
          tag: 'hr',
        },
        container: {
          type: 'group',
          schema: {
            first_name: {
              type: 'text',
              placeholder: 'First name',
              columns: {
                container: 6,
                label: 12,
                wrapper: 12,
              },
              fieldName: 'First name',
              rules: [
                'required',
                'max:255',
              ],
            },
            last_name: {
              type: 'text',
              placeholder: 'Last name',
              columns: {
                container: 6,
                label: 12,
                wrapper: 12,
              },
              fieldName: 'Last name',
              rules: [
                'required',
                'max:255',
              ],
            },
          },
          description: 'This must be your legal name.',
        },
        preferred_name_container: {
          type: 'group',
          schema: {
            preferred_name: {
              type: 'text',
              placeholder: 'Preferred name',
              columns: {
                label: 12,
                wrapper: 12,
              },
              fieldName: 'preferredName',
              rules: [
                'max:255',
              ],
            },
          },
          description: '(Optional) Your preferred name.',
        },
        age: {
          type: 'text',
          inputType: 'number',
          rules: [
            'required',
            'max:100',
          ],
          fieldName: 'age',
          placeholder: 'Age',
        },
        phone: {
          type: 'phone',
          placeholder: 'Phone',
          rules: [
            'required',
          ],
          fieldName: 'Phone',
          allowIncomplete: true,
          unmask: true,
        },
        email: {
          type: 'text',
          inputType: 'email',
          rules: [
            'required',
            'max:255',
            'email',
          ],
          placeholder: 'Email',
          fieldName: 'Email',
          description: 'School email preferred.',
        },
        country: {
          type: 'select',
          search: true,
          native: false,
          inputType: 'search',
          autocomplete: 'enabled',
          placeholder: 'Country of Residence',
          items: countriesJson,
          rules: [
            'required',
          ],
          default: 'United States',
        },
        uni: {
          type: 'select',
          search: true,
          native: false,
          inputType: 'search',
          autocomplete: 'enabled',
          placeholder: 'University',
          items: schoolsJson,
          rules: [
            'required',
          ],
          strict: false,
        },
        lvlofstudy: {
          type: 'radiogroup',
          items: [
            {
              value: 'Undergraduate University (2 year - community college or similar)',
              label: 'Undergraduate University (2 year - community college or similar)',
            },
            {
              value: 'Undergraduate University (3+ year)',
              label: 'Undergraduate University (3+ year)',
            },
            {
              value: 'Graduate University (Masters, Professional, Doctoral, etc)',
              label: 'Graduate University (Masters, Professional, Doctoral, etc)',
            },
            {
              value: 'Code School / Bootcamp',
              label: 'Code School / Bootcamp',
            },
            {
              value: 'Other Vocational / Trade Program or Apprenticeship',
              label: 'Other Vocational / Trade Program or Apprenticeship',
            },
            {
              value: 'Prefer not to answer',
              label: 'Prefer not to answer',
            },
            {
              value: 'I’m not currently a student',
              label: 'I’m not currently a student',
            },
          ],
          label: 'Highest Level of Study ',
          rules: [
            'required',
          ],
        },
        firsthack: {
          type: 'radiogroup',
          items: [
            {
              value: 'y',
              label: "Yes (We're honored to be your first hackathon experience!)",
            },
            {
              value: 'n',
              label: 'No (Welcome back!)',
            },
          ],
          label: 'Is this your first Hackathon?',
          rules: [
            'required',
          ],
        },
        major: {
          type: 'radiogroup',
          items: [
            {
              value: 'bis',
              label: 'Business & Information Systems (BIS)',
            },
            {
              value: 'cbus',
              label: 'Computing & Business (CBUS)',
            },
            {
              value: 'cs',
              label: 'Computer Science (CS)',
            },
            {
              value: 'coe',
              label: 'Computer Engineering (CoE)',
            },
            {
              value: 'ds',
              label: 'Data Science (DS)',
            },
            {
              value: 'hci',
              label: 'Human Computer Interaction (HCI)',
            },
            {
              value: 'is',
              label: 'Information Systems (IS)',
            },
            {
              value: 'it',
              label: 'Information Technology (IT)',
            },
            {
              value: 'se',
              label: 'Software Engineering (SE)',
            },
            {
              value: 'wis',
              label: 'Web & Information Systems (WIS)',
            },
            {
              value: 'othrengr',
              label: 'Another engineering discipline (such as civil, electrical, mechanical, etc.)',
            },
            {
              value: 'othrnatsci',
              label: 'A natural science (such as biology, chemistry, physics, etc.)',
            },
            {
              value: 'mathstats',
              label: 'Mathematics or Statistics',
            },
            {
              value: 'business',
              label: 'Business Discipline (such as accounting, finance, marketing, etc.)',
            },
            {
              value: 'health',
              label: 'Health Science (such as nursing, pharmacy, radiology, etc.)',
            },
            {
              value: 'arts',
              label: 'Fine arts or performing arts (such as graphic design, music, studio art, etc.)',
            },
            {
              value: 'other',
              label: 'Other',
            },
          ],
          label: 'Major',
          rules: [
            'required',
          ],
        },
        shirtsize: {
          type: 'radiogroup',
          items: [
            {
              value: 's',
              label: 'S',
            },
            {
              value: 'm',
              label: 'M',
            },
            {
              value: 'large',
              label: 'L',
            },
            {
              value: 'xl',
              label: 'XL',
            },
            {
              value: 'xxl',
              label: 'XXL',
            },
          ],
          label: 'Unisex T-Shirt Size',
          rules: [
            'required',
          ],
        },
        dietaryrestrictions: {
          type: 'checkboxgroup',
          items: [
            {
              value: '0',
              label: 'Vegetarian',
            },
            {
              value: '1',
              label: 'Vegan',
            },
            {
              value: '2',
              label: 'Kosher',
            },
            {
              value: '3',
              label: 'Halal',
            },
            {
              value: '4',
              label: 'Gluten-free',
            },
            {
              value: '5',
              label: 'None',
            },
          ],
          label: 'Dietary Restrictions ',
          rules: [
            'required',
          ],
        },
        h1: {
          type: 'static',
          tag: 'h1',
          content: 'MLH Partner Agreements',
        },
        divider_4: {
          type: 'static',
          tag: 'hr',
        },
        terms: {
          type: 'checkbox',
          text: 'I have read and agree to the <a href=https://static.mlh.io/docs/mlh-code-of-conduct.pdf>MLH Code of Conduct</a>',
          rules: [
            'accepted',
          ],
        },
        marketing_emails: {
          type: 'checkbox',
          text: 'I authorize you to share my application/registration information with Major League Hacking for event administration, ranking, and MLH administration in-line with the <a href="https://mlh.io/privacy">MLH Privacy Policy</a>. I further agree to the terms of both the <a href="https://github.com/MLH/mlh-policies/blob/main/contest-terms.md">MLH Contest Terms and Conditions</a>  and the <a href="https://mlh.io/privacy">MLH Privacy Policy</a>.',
          rules: [
            'accepted',
          ],
        },
        marketing_emails_1: {
          type: 'checkbox',
          text: '(Optional) I authorize MLH to send me occasional emails about relevant events, career opportunities, and community announcements.',
        },
        divider_3: {
          type: 'static',
          tag: 'hr',
        },
        divider_1: {
          type: 'static',
          tag: 'hr',
        },
        h1_2: {
          type: 'static',
          tag: 'h1',
          content: 'Optional Demographics',
        },
        p_1: {
          type: 'static',
          tag: 'p',
          content: '<div>This section is optional, but it helps us to get to know and serve our audience better!</div>',
        },
        minority: {
          type: 'radiogroup',
          label: 'Do you identify as part of an underrepresented group in the technology industry? ',
          items: [
            {
              value: 'y',
              label: 'Yes',
            },
            {
              value: 'n',
              label: 'No',
            },
            {
              value: 'u',
              label: 'Unsure',
            },
            {
              value: 'p',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        gender: {
          type: 'radiogroup',
          label: 'Gender',
          items: [
            {
              value: 'm',
              label: 'Male',
            },
            {
              value: 'f',
              label: 'Female',
            },
            {
              value: 'nb',
              label: 'Non-Binary',
            },
            {
              value: 'self',
              label: 'Prefer to Self-Describe',
            },
            {
              value: 'p',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        race: {
          type: 'radiogroup',
          label: 'Race / Ethnicity ',
          items: [
            {
              value: 'ai',
              label: 'Asian Indian',
            },
            {
              value: 'boa',
              label: 'Black or African',
            },
            {
              value: 'c',
              label: 'Chinese',
            },
            {
              value: 'f',
              label: 'Filipino',
            },
            {
              value: 'goc',
              label: 'Gaumanian or Chamorro',
            },
            {
              value: 'hlso',
              label: 'Hispanic / Latino / Spanish Origin',
            },
            {
              value: 'j',
              label: 'Japenese',
            },
            {
              value: 'k',
              label: 'Korean',
            },
            {
              value: 'me',
              label: 'Middle Eastern',
            },
            {
              value: 'naan',
              label: 'Native Hawaiian',
            },
            {
              value: 'nh',
              label: 'Samoan',
            },
            {
              value: 's',
              label: 'Samoan',
            },
            {
              value: 'v',
              label: 'Vietnamese',
            },
            {
              value: 'w',
              label: 'White',
            },
            {
              value: 'oa',
              label: 'Other Asian (Thai, Cambodian, etc)',
            },
            {
              value: 'opi',
              label: 'Other Pacific Islander',
            },
            {
              value: 'o',
              label: 'Other (Please Specify)',
            },
            {
              value: 'p',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        divider_6: {
          type: 'static',
          tag: 'hr',
        },
        h1_1: {
          type: 'static',
          tag: 'h1',
          content: 'Networking (Optional)',
        },
        // p: {
        //   type: 'static',
        //   tag: 'p',
        //   content: '<div>You can skip any questions in this section, but we would appreciate your response! We are working to help our sponsors and hackers connect. To that end, we will be sharing a resume booklet with our sponsors, as well as LinkedIn accounts. This information will be shared with all our sponsors.</div>',
        // },
        divider_2: {
          type: 'static',
          tag: 'hr',
        },
        // resume: {
        //   type: 'file',
        //   label: 'Resume',
        //   urls: {},
        //   description: 'Attach your resume as a PDF',
        //   accept: 'pdf,application/pdf',
        //   rules: [
        //     'max:10000',
        //   ],
        // },
        linkedin: {
          type: 'text',
          label: 'LinkedIn Account URL',
          inputType: 'url',
          placeholder: 'https://www.linkedin.com/in/YOURUSER/',
        },
        divider_5: {
          type: 'static',
          tag: 'hr',
        },
      },
    }
  })
}
</script>

<style>
.vf-registration *,
.vf-registration *:before,
.vf-registration *:after,
.vf-registration:root {
  --vf-primary: #07bf9b;
  --vf-primary-darker: #06ac8b;
  --vf-color-on-primary: #ffffff;
  --vf-danger: #ef4444;
  --vf-danger-lighter: #fee2e2;
  --vf-success: #10b981;
  --vf-success-lighter: #d1fae5;
  --vf-gray-50: #f9fafb;
  --vf-gray-100: #f3f4f6;
  --vf-gray-200: #e5e7eb;
  --vf-gray-300: #d1d5db;
  --vf-gray-400: #9ca3af;
  --vf-gray-500: #6b7280;
  --vf-gray-600: #4b5563;
  --vf-gray-700: #374151;
  --vf-gray-800: #1f2937;
  --vf-gray-900: #111827;
  --vf-dark-50: #EFEFEF;
  --vf-dark-100: #DCDCDC;
  --vf-dark-200: #BDBDBD;
  --vf-dark-300: #A0A0A0;
  --vf-dark-400: #848484;
  --vf-dark-500: #737373;
  --vf-dark-600: #393939;
  --vf-dark-700: #323232;
  --vf-dark-800: #262626;
  --vf-dark-900: #191919;
  --vf-ring-width: 2px;
  --vf-ring-color: #07bf9b66;
  --vf-link-color: var(--vf-primary);
  --vf-link-decoration: inherit;
  --vf-font-size: 1rem;
  --vf-font-size-sm: 0.875rem;
  --vf-font-size-lg: 1rem;
  --vf-font-size-small: 0.875rem;
  --vf-font-size-small-sm: 0.8125rem;
  --vf-font-size-small-lg: 0.875rem;
  --vf-font-size-h1: 2.125rem;
  --vf-font-size-h1-sm: 2.125rem;
  --vf-font-size-h1-lg: 2.125rem;
  --vf-font-size-h2: 1.875rem;
  --vf-font-size-h2-sm: 1.875rem;
  --vf-font-size-h2-lg: 1.875rem;
  --vf-font-size-h3: 1.5rem;
  --vf-font-size-h3-sm: 1.5rem;
  --vf-font-size-h3-lg: 1.5rem;
  --vf-font-size-h4: 1.25rem;
  --vf-font-size-h4-sm: 1.25rem;
  --vf-font-size-h4-lg: 1.25rem;
  --vf-font-size-h1-mobile: 1.5rem;
  --vf-font-size-h1-mobile-sm: 1.5rem;
  --vf-font-size-h1-mobile-lg: 1.5rem;
  --vf-font-size-h2-mobile: 1.25rem;
  --vf-font-size-h2-mobile-sm: 1.25rem;
  --vf-font-size-h2-mobile-lg: 1.25rem;
  --vf-font-size-h3-mobile: 1.125rem;
  --vf-font-size-h3-mobile-sm: 1.125rem;
  --vf-font-size-h3-mobile-lg: 1.125rem;
  --vf-font-size-h4-mobile: 1rem;
  --vf-font-size-h4-mobile-sm: 1rem;
  --vf-font-size-h4-mobile-lg: 1rem;
  --vf-font-size-blockquote: 1rem;
  --vf-font-size-blockquote-sm: 0.875rem;
  --vf-font-size-blockquote-lg: 1rem;
  --vf-line-height: 1.5rem;
  --vf-line-height-sm: 1.25rem;
  --vf-line-height-lg: 1.5rem;
  --vf-line-height-small: 1.25rem;
  --vf-line-height-small-sm: 1.125rem;
  --vf-line-height-small-lg: 1.25rem;
  --vf-line-height-headings: 1.2;
  --vf-line-height-headings-sm: 1.2;
  --vf-line-height-headings-lg: 1.2;
  --vf-line-height-blockquote: 1.5rem;
  --vf-line-height-blockquote-sm: 1.25rem;
  --vf-line-height-blockquote-lg: 1.5rem;
  --vf-letter-spacing: 0px;
  --vf-letter-spacing-sm: 0px;
  --vf-letter-spacing-lg: 0px;
  --vf-letter-spacing-small: 0px;
  --vf-letter-spacing-small-sm: 0px;
  --vf-letter-spacing-small-lg: 0px;
  --vf-letter-spacing-headings: 0px;
  --vf-letter-spacing-headings-sm: 0px;
  --vf-letter-spacing-headings-lg: 0px;
  --vf-letter-spacing-blockquote: 0px;
  --vf-letter-spacing-blockquote-sm: 0px;
  --vf-letter-spacing-blockquote-lg: 0px;
  --vf-gutter: 1rem;
  --vf-gutter-sm: 0.5rem;
  --vf-gutter-lg: 1rem;
  --vf-min-height-input: 2.375rem;
  --vf-min-height-input-sm: 2.125rem;
  --vf-min-height-input-lg: 2.875rem;
  --vf-py-input: 0.375rem;
  --vf-py-input-sm: 0.375rem;
  --vf-py-input-lg: 0.625rem;
  --vf-px-input: 0.75rem;
  --vf-px-input-sm: 0.5rem;
  --vf-px-input-lg: 0.875rem;
  --vf-py-btn: 0.375rem;
  --vf-py-btn-sm: 0.375rem;
  --vf-py-btn-lg: 0.625rem;
  --vf-px-btn: 0.875rem;
  --vf-px-btn-sm: 0.75rem;
  --vf-px-btn-lg: 1.25rem;
  --vf-py-btn-small: 0.25rem;
  --vf-py-btn-small-sm: 0.25rem;
  --vf-py-btn-small-lg: 0.375rem;
  --vf-px-btn-small: 0.625rem;
  --vf-px-btn-small-sm: 0.625rem;
  --vf-px-btn-small-lg: 0.75rem;
  --vf-py-group-tabs: 0.375rem;
  --vf-py-group-tabs-sm: 0.375rem;
  --vf-py-group-tabs-lg: 0.625rem;
  --vf-px-group-tabs: 0.75rem;
  --vf-px-group-tabs-sm: 0.5rem;
  --vf-px-group-tabs-lg: 0.875rem;
  --vf-py-group-blocks: 0.75rem;
  --vf-py-group-blocks-sm: 0.625rem;
  --vf-py-group-blocks-lg: 0.875rem;
  --vf-px-group-blocks: 1rem;
  --vf-px-group-blocks-sm: 1rem;
  --vf-px-group-blocks-lg: 1rem;
  --vf-py-tag: 0px;
  --vf-py-tag-sm: 0px;
  --vf-py-tag-lg: 0px;
  --vf-px-tag: 0.4375rem;
  --vf-px-tag-sm: 0.4375rem;
  --vf-px-tag-lg: 0.4375rem;
  --vf-py-slider-tooltip: 0.125rem;
  --vf-py-slider-tooltip-sm: 0.0625rem;
  --vf-py-slider-tooltip-lg: 0.1875rem;
  --vf-px-slider-tooltip: 0.375rem;
  --vf-px-slider-tooltip-sm: 0.3125rem;
  --vf-px-slider-tooltip-lg: 0.5rem;
  --vf-py-blockquote: 0.25rem;
  --vf-py-blockquote-sm: 0.25rem;
  --vf-py-blockquote-lg: 0.25rem;
  --vf-px-blockquote: 0.75rem;
  --vf-px-blockquote-sm: 0.75rem;
  --vf-px-blockquote-lg: 0.75rem;
  --vf-py-hr: 0.25rem;
  --vf-space-addon: 0px;
  --vf-space-addon-sm: 0px;
  --vf-space-addon-lg: 0px;
  --vf-space-checkbox: 0.375rem;
  --vf-space-checkbox-sm: 0.375rem;
  --vf-space-checkbox-lg: 0.375rem;
  --vf-space-tags: 0.1875rem;
  --vf-space-tags-sm: 0.1875rem;
  --vf-space-tags-lg: 0.1875rem;
  --vf-space-static-tag-1: 1rem;
  --vf-space-static-tag-2: 2rem;
  --vf-space-static-tag-3: 3rem;
  --vf-floating-top: 0rem;
  --vf-floating-top-sm: 0rem;
  --vf-floating-top-lg: 0.6875rem;
  --vf-bg-input: #ffffff;
  --vf-bg-input-hover: #ffffff;
  --vf-bg-input-focus: #ffffff;
  --vf-bg-input-danger: #ffffff;
  --vf-bg-input-success: #ffffff;
  --vf-bg-checkbox: #ffffff;
  --vf-bg-checkbox-hover: #ffffff;
  --vf-bg-checkbox-focus: #ffffff;
  --vf-bg-checkbox-danger: #ffffff;
  --vf-bg-checkbox-success: #ffffff;
  --vf-bg-disabled: var(--vf-gray-200);
  --vf-bg-selected: #1118270d;
  --vf-bg-passive: var(--vf-gray-300);
  --vf-bg-icon: var(--vf-gray-500);
  --vf-bg-danger: var(--vf-danger-lighter);
  --vf-bg-success: var(--vf-success-lighter);
  --vf-bg-tag: var(--vf-primary);
  --vf-bg-slider-handle: var(--vf-primary);
  --vf-bg-toggle-handle: #ffffff;
  --vf-bg-date-head: var(--vf-gray-100);
  --vf-bg-addon: #ffffff00;
  --vf-bg-btn: var(--vf-primary);
  --vf-bg-btn-danger: var(--vf-danger);
  --vf-bg-btn-secondary: var(--vf-gray-200);
  --vf-bg-table-header: var(--vf-gray-100);
  --vf-color-input: var(--vf-gray-800);
  --vf-color-input-hover: var(--vf-gray-800);
  --vf-color-input-focus: var(--vf-gray-800);
  --vf-color-input-danger: var(--vf-gray-800);
  --vf-color-input-success: var(--vf-gray-800);
  --vf-color-disabled: var(--vf-gray-400);
  --vf-color-placeholder: var(--vf-gray-300);
  --vf-color-passive: var(--vf-gray-700);
  --vf-color-muted: var(--vf-gray-500);
  --vf-color-floating: var(--vf-gray-500);
  --vf-color-floating-focus: var(--vf-gray-500);
  --vf-color-floating-success: var(--vf-gray-500);
  --vf-color-floating-danger: var(--vf-gray-500);
  --vf-color-danger: var(--vf-danger);
  --vf-color-success: var(--vf-success);
  --vf-color-tag: var(--vf-color-on-primary);
  --vf-color-addon: var(--vf-gray-800);
  --vf-color-date-head: var(--vf-gray-700);
  --vf-color-btn: var(--vf-color-on-primary);
  --vf-color-btn-danger: #ffffff;
  --vf-color-btn-secondary: var(--vf-gray-700);
  --vf-color-table-header: var(--vf-gray-800);
  --vf-border-color-input: var(--vf-gray-300);
  --vf-border-color-input-hover: var(--vf-gray-300);
  --vf-border-color-input-focus: var(--vf-primary);
  --vf-border-color-input-danger: var(--vf-gray-300);
  --vf-border-color-input-success: var(--vf-gray-300);
  --vf-border-color-checkbox: var(--vf-gray-300);
  --vf-border-color-checkbox-focus: var(--vf-primary);
  --vf-border-color-checkbox-hover: var(--vf-gray-300);
  --vf-border-color-checkbox-danger: var(--vf-gray-300);
  --vf-border-color-checkbox-success: var(--vf-gray-300);
  --vf-border-color-checked: var(--vf-primary);
  --vf-border-color-passive: var(--vf-gray-300);
  --vf-border-color-slider-tooltip: var(--vf-primary);
  --vf-border-color-tag: var(--vf-primary);
  --vf-border-color-btn: var(--vf-primary);
  --vf-border-color-btn-danger: var(--vf-danger);
  --vf-border-color-btn-secondary: var(--vf-gray-200);
  --vf-border-color-blockquote: var(--vf-gray-300);
  --vf-border-color-hr: var(--vf-gray-300);
  --vf-border-color-signature-hr: var(--vf-gray-300);
  --vf-border-color-table: var(--vf-gray-300);
  --vf-border-width-input-t: 1px;
  --vf-border-width-input-r: 1px;
  --vf-border-width-input-b: 1px;
  --vf-border-width-input-l: 1px;
  --vf-border-width-radio-t: 1px;
  --vf-border-width-radio-r: 1px;
  --vf-border-width-radio-b: 1px;
  --vf-border-width-radio-l: 1px;
  --vf-border-width-checkbox-t: 1px;
  --vf-border-width-checkbox-r: 1px;
  --vf-border-width-checkbox-b: 1px;
  --vf-border-width-checkbox-l: 1px;
  --vf-border-width-dropdown: 1px;
  --vf-border-width-btn: 1px;
  --vf-border-width-toggle: 0.125rem;
  --vf-border-width-tag: 1px;
  --vf-border-width-blockquote: 3px;
  --vf-border-width-table: 1px;
  --vf-shadow-input: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-input-hover: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-input-focus: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-handles: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-handles-hover: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-handles-focus: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-btn: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-shadow-dropdown: 0px 0px 0px 0px rgba(0,0,0,0);
  --vf-radius-input: 0.25rem;
  --vf-radius-input-sm: 0.25rem;
  --vf-radius-input-lg: 0.25rem;
  --vf-radius-btn: 0.25rem;
  --vf-radius-btn-sm: 0.25rem;
  --vf-radius-btn-lg: 0.25rem;
  --vf-radius-small: 0.25rem;
  --vf-radius-small-sm: 0.25rem;
  --vf-radius-small-lg: 0.25rem;
  --vf-radius-large: 0.25rem;
  --vf-radius-large-sm: 0.25rem;
  --vf-radius-large-lg: 0.25rem;
  --vf-radius-tag: 0.25rem;
  --vf-radius-tag-sm: 0.25rem;
  --vf-radius-tag-lg: 0.25rem;
  --vf-radius-checkbox: 0.25rem;
  --vf-radius-checkbox-sm: 0.25rem;
  --vf-radius-checkbox-lg: 0.25rem;
  --vf-radius-slider: 0.25rem;
  --vf-radius-slider-sm: 0.25rem;
  --vf-radius-slider-lg: 0.25rem;
  --vf-radius-image: 0.25rem;
  --vf-radius-image-sm: 0.25rem;
  --vf-radius-image-lg: 0.25rem;
  --vf-radius-gallery: 0.25rem;
  --vf-radius-gallery-sm: 0.25rem;
  --vf-radius-gallery-lg: 0.25rem;
  --vf-checkbox-size: 1rem;
  --vf-checkbox-size-sm: 0.875rem;
  --vf-checkbox-size-lg: 1rem;
  --vf-gallery-size: 6rem;
  --vf-gallery-size-sm: 5rem;
  --vf-gallery-size-lg: 7rem;
  --vf-toggle-width: 3rem;
  --vf-toggle-width-sm: 2.75rem;
  --vf-toggle-width-lg: 3rem;
  --vf-toggle-height: 1.25rem;
  --vf-toggle-height-sm: 1rem;
  --vf-toggle-height-lg: 1.25rem;
  --vf-slider-height: 0.375rem;
  --vf-slider-height-sm: 0.3125rem;
  --vf-slider-height-lg: 0.5rem;
  --vf-slider-height-vertical: 20rem;
  --vf-slider-height-vertical-sm: 20rem;
  --vf-slider-height-vertical-lg: 20rem;
  --vf-slider-handle-size: 1rem;
  --vf-slider-handle-size-sm: 0.875rem;
  --vf-slider-handle-size-lg: 1.25rem;
  --vf-slider-tooltip-distance: 0.5rem;
  --vf-slider-tooltip-distance-sm: 0.375rem;
  --vf-slider-tooltip-distance-lg: 0.5rem;
  --vf-slider-tooltip-arrow-size: 0.3125rem;
  --vf-slider-tooltip-arrow-size-sm: 0.3125rem;
  --vf-slider-tooltip-arrow-size-lg: 0.3125rem;
}

.vf-description {
  text-align: left;
}

*[aria-labelledby="container__label"] {
  background-color: #eef3f7;
  padding: 8px;
  border-radius: 8px;
}

</style>
